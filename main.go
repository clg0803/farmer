package main

import (
	"farmer/impl"
	"fmt"
	"io"
	"net"
)

func handle(conn net.Conn) {
	pker := impl.NewPacker()
	for {
		headData := make([]byte, pker.HeaderLen())
		_, err := io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println(":[ERR]: READ HEADER ERR", err)
			break
		}
		msg, err := pker.Unpack(headData)
		if err != nil {
			fmt.Println(":[ERR]: UNPACK HEADER ERR", err)
			break
		}
		if msg.DataLen() > 0 {
			msg.SetData(make([]byte, msg.DataLen()))
			_, err := io.ReadFull(conn, msg.Data())
			if err != nil {
				fmt.Println(":[ERR]: UNPACK DATA ERR", err)
				break
			}
			fmt.Printf("=> RECV MSG : ID = %d\n, LEN = %d\n, DATA = %s\n",
				msg.MsgID(),
				msg.DataLen(),
				string(msg.Data()),
			)
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println(":[ERR]: SERVER LISTENER NOT ESTABLISHED ", err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(":[ERR]: SERVER ACCEPT ERR ", err)
			return
		}
		go handle(conn)
	}
}
