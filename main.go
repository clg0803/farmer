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

func (p *PingRouter) Handle(req iface.IRequest) {
	fmt.Println("CALL OF HANDLE_FUNCTION")
	_, err := req.GetConnection().GetTcpConnection().Write([]byte("handler_function ... "))
	if err != nil {
		fmt.Println(":[ERR]: HANDLE_FUNCTION")
	}
}

func (p *PingRouter) After(req iface.IRequest) {
	fmt.Println("CALL OF AFTER_FUNCTION")
	_, err := req.GetConnection().GetTcpConnection().Write([]byte("after_function ... "))
	if err != nil {
		fmt.Println(":[ERR]: AFTER_FUNCTION")
	}
}

func main() {
	s := impl.NewServer("[farmer v0.3]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
