package impl

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func Client() {
	for i := 0; i < 5; i++ {
		fmt.Println("<Client> connect ... ")
		time.Sleep(1 * time.Second)
		conn, err := net.Dial("tcp", "127.0.0.1:8848")
		if err != nil {
			fmt.Println(":<ERR>: ERR CREATE CONN ", err)
			return
		}
		_, err = conn.Write([]byte("Hello Framer"))
		if err != nil {
			fmt.Println(":<ERR>: ERR WRITE TO SERVER ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println(":<ERR>: ERR RECEIVE FROM SERVER ", err)
			return
		}
		fmt.Printf(":<SUCCESS>: SERVER CALL BACK: %s, CNT = %d \n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}

func TestSever(t *testing.T) {
	s := NewServer("[farmer v0.2]")
	go Client()
	s.Serve() // don't end
}
