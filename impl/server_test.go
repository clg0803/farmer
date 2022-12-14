package impl

import (
	"fmt"
	"net"
	"testing"
	"time"
)

// implement CLIENT

func Client() {
	fmt.Println("<Client> TEST START ... ")
	time.Sleep(2 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println(":<ERR>: ERR CREATE CONN ", err)
		return
	}
	for {
		packer := NewPacker()
		data, _ := packer.Pack(NewMessage(uint32(1), []byte("Client try to connect")))
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println(":<ERR>: CLIENT WRITE ERR")
			return
		}
		msg, err := packer.ReadAndUnpackToMsg(NewConnection(nil, conn.(*net.TCPConn), 1001, nil))
		if err != nil {
			fmt.Println(":<ERR>: COVERT DATA TO MSG ERR", err)
		}
		fmt.Printf(":<SUCCESS>: RECV FROM SERVER, MSG ID = %d,\n DATA = %s\n",
			msg.MsgID(), string(msg.Data()),
		)
		time.Sleep(1 * time.Second)
	}
	conn.Close()
}

func TestSever(t *testing.T) {
	Client()
}
