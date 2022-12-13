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
	err := r.GetConnection().SendMsg(r.GetMsgID()+100, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(":[ERR]: WRITE BACK ERR", err)
	}
}

func main() {
	s := impl.NewServer("[farmer v0.5]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
