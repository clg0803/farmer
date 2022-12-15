package iface

import "net"

// to bind a connection with a func()

type IConnection interface {
	Start()
	Stop()
	GetTcpConnection() *net.TCPConn // get TCP connection bind with it
	GetConnectionID() uint32
	GetRemoteAddr() net.Addr                         // get remote(client) addr
	SendMsg(msgID uint32, data []byte) error         // send MSG to CLIENT
	SendMsgWithBuff(msgID uint32, data []byte) error // send MSG to CLIENT

	SetProperty(key string, value interface{})
	GetProperty(key string) (interface{}, error)
	DelProperty(key string)
}

type ToHandle func(conn *net.TCPConn, data []byte, length int) error //Client register ToHandle() with an IConnection
