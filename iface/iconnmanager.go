package iface

type IConnectManager interface {
	Add(conn IConnection)
	Remove(conn IConnection)
	Get(connID uint32) (IConnection, error)
	ConnectedNum() int32
	CleanAllConn()
}
