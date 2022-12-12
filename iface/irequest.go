package iface

// IRequest
// contains Connection and associated data
type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
}
