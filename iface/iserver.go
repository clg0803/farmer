package iface

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(msgId uint32, r IRouter)
	GetConnMgr() IConnectManager
	SetOnConnStart(func(connection IConnection))
	SetOnConnEnd(func(connection IConnection))
	CallOnConnStart(c IConnection)
	CallOnConnEnd(c IConnection)
}
