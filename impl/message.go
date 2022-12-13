package impl

type Message struct {
	id      uint32
	data    []byte
	dataLen uint32
}

func (m *Message) Data() []byte {
	return m.data
}

func (m *Message) DataLen() uint32 {
	return m.dataLen
}

func (m *Message) MsgID() uint32 {
	return m.id
}

func (m *Message) SetData(data []byte) {
	m.data = data
}

func (m *Message) SetDataLen(len uint32) {
	m.dataLen = len
}

func (m *Message) SetMsgID(id uint32) {
	m.id = id
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		id:      id,
		data:    data,
		dataLen: uint32(len(data)),
	}
}
