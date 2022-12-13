package iface

// IPack
// define pack/unpack ways
// pack format is defined in `imessage`

//	[	  HEAD   ][data]
//
// <--> length | id | data <-->
type IPacker interface {
	HeaderLen() uint32
	Pack(msg IMessage) ([]byte, error)                      // Pack IMessage to byte stream to transmit
	Unpack([]byte) (IMessage, error)                        // Convert byte stream to IMessage obj
	ReadAndUnpackToMsg(conn *IConnection) (IMessage, error) // read []byte from Conn and convert to IMessage
}
