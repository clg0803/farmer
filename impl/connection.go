package impl

import (
	"farmer/iface"
	"net"
)

type Connection struct {
	conn     net.TCPConn
	connID   uint32
	isClosed bool
	exitChan chan bool
	farm     iface.ToHandle
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
