package impl

import (
	"farmer/iface"
)

type Request struct {
	conn iface.IConnection
	msg  iface.IMessage
}

func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

func (r *Request) GetMsgData() []byte {
	return r.msg.Data()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.MsgID()
}
