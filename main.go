package main

import (
	"farmer/iface"
	"farmer/impl"
	"fmt"
)

// PingRouter a simple subclass of BaseRouter
type PingRouter struct {
	impl.BaseRouter
}

func (p *PingRouter) Before(req iface.IRequest) {
	fmt.Println("CALL OF BEFORE_FUNCTION")
	_, err := req.GetConnection().GetTcpConnection().Write([]byte("before_function ... "))
	if err != nil {
		fmt.Println(":[ERR]: BEFORE_FUNCTION")
	}
}

func main() {
	s := impl.NewServer("farmer0.4")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
