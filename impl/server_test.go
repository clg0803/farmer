package impl

import (
	"fmt"
	"net"
	"testing"
	"time"
)

// implement CLIENT

func Client() {
	for i := 0; i < 2; i++ {
		fmt.Println("<Client> connect ... ")
		time.Sleep(1 * time.Second)
		conn, err := net.Dial("tcp", "127.0.0.1:8848")
		if err != nil {
			fmt.Println(":<ERR>: ERR CREATE CONN ", err)
			return
		}

		packer := NewPacker()
		data, _ := packer.Pack(NewMessage(uint32(i), []byte("Client try to connect")))
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println(":<ERR>: CLIENT WRITE ERR")
			return
		}
		msg, err := packer.ReadAndUnpackToMsg(NewConnection(conn.(*net.TCPConn), 1001, nil))
		if err != nil {
			fmt.Println(":<ERR>: COVERT DATA TO MSG ERR", err)
			continue
		}
		fmt.Printf(":<SUCCESS>: RECV FROM SERVER, MSG ID = %d,\n DATA = %s\n",
			msg.MsgID(), string(msg.Data()),
		)
	}
}

func TestSever(t *testing.T) {
	Client()
}
