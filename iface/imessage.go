package iface

// IMessage
// wrap bare data
// a typical 'farmer' pack contains:
//
//		 [	  HEAD   ][data]
//	<--> length | id | data <-->
//
// we cant specific data by HEAD(length, id), split by `length`
type IMessage interface {
	Data() []byte
	DataLen() uint32
	MsgID() uint32

	SetData([]byte)
	SetDataLen(uint32)
	SetMsgID(uint32)
}
