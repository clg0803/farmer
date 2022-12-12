package impl

import (
	"farmer/iface"
	"fmt"
	"net"
)

type Connection struct {
	conn     net.TCPConn
	connID   uint32
	isClosed bool
	exitChan chan bool
	farm     iface.ToHandle
}

// Start -- work for each connection
// use `go c.Start()` to invoke
func (c *Connection) Start() {
	fmt.Println(":[START]: CONN_ID = ", c.connID)
	c.readAndHandle()
	c.Send(nil)
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

func (c *Connection) Send(data []byte) error {
	return nil
}

// read from client and 'farm' it
func (c *Connection) readAndHandle() {
	fmt.Println(":[SUCCESS]: START READING ...")
	defer c.Stop() // in case of failing
	defer fmt.Println(":[SUCCESS]: stop READING from ", c.GetRemoteAddr().String())
	for {
		buf := make([]byte, 512)
		cnt, err := c.conn.Read(buf)
		if err != nil {
			fmt.Println(":[ERR]: READ ERR", err)
			continue
		}
		// invoke `farm()` to handle
		if err := c.farm(&c.conn, buf, cnt); err != nil {
			fmt.Println(":[ERR]: FAILED TO HANDLE ", c.connID)
			fmt.Println(err)
			break
		}
	}
}

func NewConnection(tcpConn *net.TCPConn, connID uint32, toHandle iface.ToHandle) *Connection {
	return &Connection{
		conn:   *tcpConn,
		connID: connID,
		farm:   toHandle,

		isClosed: false,
		exitChan: make(chan bool, 1),
	}
}
