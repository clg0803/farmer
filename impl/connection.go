package impl

import (
	"errors"
	"farmer/iface"
	"fmt"
	"net"
)

type Connection struct {
	server   iface.IServer
	conn     net.TCPConn
	connID   uint32
	isClosed bool
	exitChan chan bool

	msgHandler  iface.IMsgHandler // for R/W separated
	msgChan     chan []byte
	msgBuffChan chan []byte
}

// Start -- work for each connection
// use `go c.Start()` to invoke
func (c *Connection) Start() {
	fmt.Println(":[START]: CONN_ID = ", c.connID)
	go c.readAndHandle()
	go c.startWriter()
	c.server.CallOnConnStart(c)
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
	c.server.CallOnConnEnd(c)
	c.conn.Close()
	c.exitChan <- true
	c.server.GetConnMgr().Remove(c)
	close(c.msgChan)
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
	// write to msgChan, let writer do
	c.msgChan <- msg
	return nil
}

func (c *Connection) SendMsgWithBuff(msgID uint32, data []byte) error {
	if c.isClosed {
		return errors.New(":[ERR]: CONNECTION CLOSED BEFORE SENDING MSG")
	}
	pker := NewPacker()
	msg, err := pker.Pack(NewMessage(msgID, data))
	if err != nil {
		fmt.Println(":[ERR]: PACK ERR, MSG ID = ", msgID)
		return errors.New(":[ERR]: PACK ERR")
	}
	// write to msgBuffChan, let writer do
	c.msgBuffChan <- msg
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
		c.msgHandler.SendMsgToTaskQueue(&req)
	}
}

func (c *Connection) startWriter() {
	fmt.Println(":[SUCCESS]: CO-WRITER START ...")
	defer fmt.Println(":[SUCCESS]: WRITING TO <", c.GetRemoteAddr(), "> FINISHED, WRITER EXITS")
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.conn.Write(data); err != nil {
				fmt.Println(":[ERR]: WRITING TO CLIENT ERR, ", err)
				return
			}
		case data, ok := <-c.msgBuffChan:
			if ok {
				if _, err := c.conn.Write(data); err != nil {
					fmt.Println(":[ERR]: WRITING TO CLIENT ERR, ", err)
					return
				}
			} else {
				fmt.Println(":[SUCCESS]: MSG_BUFF_CHAN IS CLOSED")
				break
			}
		case <-c.exitChan:
			return
		}
	}
}

func NewConnection(s iface.IServer, tcpConn *net.TCPConn, connID uint32, msgHandler iface.IMsgHandler) *Connection {
	c := Connection{
		server:      s,
		conn:        *tcpConn,
		connID:      connID,
		msgHandler:  msgHandler,
		msgChan:     make(chan []byte),
		msgBuffChan: make(chan []byte, 1024),

		isClosed: false,
		exitChan: make(chan bool, 1),
	}
	c.server.GetConnMgr().Add(&c)
	return &c
}
