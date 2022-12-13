package impl

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func Client() {
	for i := 0; i < 1; i++ {
		fmt.Println("<Client> connect ... ")
		time.Sleep(1 * time.Second)
		conn, err := net.Dial("tcp", "127.0.0.1:8848")
		if err != nil {
			fmt.Println(":<ERR>: ERR CREATE CONN ", err)
			return
		}
		packer := NewPacker()
		msg1 := &Message{
			id:      0,
			dataLen: 5,
			data:    []byte{'h', 'e', 'l', 'l', 'o'},
		}
		ds1, err := packer.Pack(msg1)
		if err != nil {
			fmt.Println(":<ERR>: CLIENT PACKING ERR", err)
			return
		}
		msg2 := &Message{
			id:      1,
			dataLen: 4,
			data:    []byte{'n', 'i', 'h', 'o'},
		}
		ds2, err := packer.Pack(msg2)
		if err != nil {
			fmt.Println(":<ERR>: CLIENT PACKING ERR", err)
			return
		}

		ds := append(ds1, ds2...)
		conn.Write(ds)
		time.Sleep(1 * time.Second)
	}
}

func TestSever(t *testing.T) {
	Client()
}
