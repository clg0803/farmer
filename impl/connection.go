package impl

import (
	"errors"
	"farmer/iface"
	"fmt"
	"net"
)

type Connection struct {
	conn     net.TCPConn
	connID   uint32
	isClosed bool
	exitChan chan bool

	msgHandler iface.IMsgHandler
}

// Start -- work for each connection
// use `go c.Start()` to invoke
func (c *Connection) Start() {
	fmt.Println(":[START]: CONN_ID = ", c.connID)
	c.readAndHandle()
	for {
		select {
		case <-c.exitChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	fmt.Println(":[STOP]: CONN_ID = ", c.connID)
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.conn.Close()
	close(c.exitChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return &c.conn
}

func (c *Connection) GetConnectionID() uint32 {
	return c.connID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.isClosed {
		return errors.New(":[ERR]: CONNECTION CLOSED BEFORE SENDING MSG")
	}
	pker := NewPacker()
	msg, err := pker.Pack(NewMessage(msgID, data))
	if err != nil {
		fmt.Println(":[ERR]: PACK ERR, MSG ID = ", msgID)
		return errors.New(":[ERR]: PACK ERR")
	}
	if _, err := c.conn.Write(msg); err != nil {
		fmt.Println(":[ERR]: SEND BACK MSG ERR, MSG ID = ", msgID)
		c.exitChan <- true
		return errors.New(":[ERR]: SEND BACK ERR")
	}
	return nil
}

// read from client and 'farm' it
func (c *Connection) readAndHandle() {
	fmt.Println(":[SUCCESS]: START READING ...")
	defer c.Stop() // in case of failing
	defer fmt.Println(":[SUCCESS]: STOP READING FROM ", c.GetRemoteAddr().String())
	for {
		pker := NewPacker()
		msg, err := pker.ReadAndUnpackToMsg(c)
		if err != nil {
			fmt.Println(":[ERR]: COVERT DATA TO MSG ERR", err)
			continue
		}
		req := Request{
			conn: c,
			msg:  msg,
		}
		go c.msgHandler.HandleRequest(&req)
	}
}

func NewConnection(tcpConn *net.TCPConn, connID uint32, msgHandler iface.IMsgHandler) *Connection {
	return &Connection{
		conn:       *tcpConn,
		connID:     connID,
		msgHandler: msgHandler,

		isClosed: false,
		exitChan: make(chan bool, 1),
	}
}
