package impl

import (
	"bytes"
	"encoding/binary"
	"errors"
	"farmer/iface"
	"fmt"
	"io"
)

var MaxPackSize uint32 = 65535

//	[	  HEAD   ][data]
//
// <--> length | id | data <-->
type Packer struct{}

func (p *Packer) HeaderLen() uint32 {
	return 8 // length(unit32) + id(uint32)
}

func (p *Packer) Pack(msg iface.IMessage) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	// write length
	if err := binary.Write(buff, binary.LittleEndian, msg.DataLen()); err != nil {
		return nil, err
	}
	// write id
	if err := binary.Write(buff, binary.LittleEndian, msg.MsgID()); err != nil {
		return nil, err
	}
	// write data
	if err := binary.Write(buff, binary.LittleEndian, msg.Data()); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func (p *Packer) Unpack(rd []byte) (iface.IMessage, error) {
	buff := bytes.NewReader(rd)
	msg := &Message{}
	// read length
	if err := binary.Read(buff, binary.LittleEndian, &msg.dataLen); err != nil {
		return nil, err
	}
	// read id
	if err := binary.Read(buff, binary.LittleEndian, &msg.id); err != nil {
		return nil, err
	}
	if msg.dataLen > MaxPackSize {
		return nil, errors.New(":[ERR]: TWO LARGE MSG DATA RECEIVED")
	}
	return msg, nil
}

func (p *Packer) ReadAndUnpackToMsg(c *Connection) (iface.IMessage, error) {
	headData := make([]byte, p.HeaderLen())
	if _, err := io.ReadFull(c.GetTcpConnection(), headData); err != nil {
		fmt.Println(":[ERR]: READ HEADER ERR", err)
		c.exitChan <- true
		return nil, err
	}
	msg, err := p.Unpack(headData)
	if err != nil {
		fmt.Println(":[ERR]: UNPACK HEADER ERR", err)
		c.exitChan <- true
		return nil, err
	}
	if msg.DataLen() > 0 {
		msg.SetData(make([]byte, msg.DataLen()))
		if _, err := io.ReadFull(c.GetTcpConnection(), msg.Data()); err != nil {
			fmt.Println(":[ERR]: UNPACK DATA ERR", err)
			c.exitChan <- true
			return nil, err
		}
	}
	return msg, nil
}

func NewPacker() *Packer {
	return &Packer{}
}
