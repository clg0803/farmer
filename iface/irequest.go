package iface

// IRequest
// contains Connection and associated data
type IRequest interface {
	GetConnection() IConnection
	GetMsgData() []byte
	GetMsgID() uint32
}
