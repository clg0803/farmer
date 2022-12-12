package iface

import "net"

// to bind a connection with a func()

type IConnection interface {
	Start()
	Stop()
	GetTcpConnection() *net.TCPConn // get TCP connection bind with it
	GetConnectionID() uint32
	GetRemoteAddr() net.Addr // get remote(client) addr
	Send(data []byte) error
}

type ToHandle func(conn *net.TCPConn, data []byte, length int) error //Client register ToHandle() with an IConnection
