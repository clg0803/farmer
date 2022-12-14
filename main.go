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
	err := r.GetConnection().SendMsg(111, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(":[ERR]: WRITE BACK ERR", err)
	}
}

type HelloRouter struct {
	impl.BaseRouter
}

func (hr *HelloRouter) Handle(r iface.IRequest) {
	fmt.Println(":[SUCCESS]: CALL OF PING_ROUTER HANDLER")
	fmt.Printf(":[SUCCESS]: RECV FROM CLIENT, MSG ID = %d,\n DATA = %s\n",
		r.GetMsgID(), string(r.GetMsgData()),
	)
	err := r.GetConnection().SendMsg(999, []byte("Pong Pong Pong"))
	if err != nil {
		fmt.Println(":[ERR]: WRITE BACK ERR", err)
	}
}

func main() {
	s := impl.NewServer("[farmer v0.6]")
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
