package main

import (
	"farmer/iface"
	"farmer/impl"
	"fmt"
)

type PingRouter struct {
	impl.BaseRouter
}

func (pr *PingRouter) Handle(r iface.IRequest) {
	fmt.Println(":[SUCCESS]: CALL OF PING_ROUTER HANDLER")
	fmt.Printf(":[SUCCESS]: RECV FROM CLIENT, MSG ID = %d,\n DATA = %s\n",
		r.GetMsgID(), string(r.GetMsgData()),
	)
	err := r.GetConnection().SendMsgWithBuff(111, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(":[ERR]: WRITE BACK ERR", err)
	}
}

type HelloRouter struct {
	impl.BaseRouter
}

func (hr *HelloRouter) Handle(r iface.IRequest) {
	fmt.Println(":[SUCCESS]: CALL OF PONG_PONG_ROUTER HANDLER")
	fmt.Printf(":[SUCCESS]: RECV FROM CLIENT, MSG ID = %d,\n DATA = %s\n",
		r.GetMsgID(), string(r.GetMsgData()),
	)
	err := r.GetConnection().SendMsgWithBuff(999, []byte("Pong Pong Pong"))
	if err != nil {
		fmt.Println(":[ERR]: WRITE BACK ERR", err)
	}
}

func onConnStart(conn iface.IConnection) {
	fmt.Println(":[HOOK]: `onConnStart` IS CALLED")
	if err := conn.SendMsg(666, []byte("RISE UP")); err != nil {
		println(err)
	}
}

func onConnEnd(conn iface.IConnection) {
	fmt.Println(":[HOOK]: `onConnEnd` IS CALLED")
}

func main() {
	s := impl.NewServer("[farmer v0.6]")

	s.SetOnConnStart(onConnStart)
	s.SetOnConnEnd(onConnEnd)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	s.Serve()
}
