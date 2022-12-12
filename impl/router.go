package impl

import (
	"farmer/iface"
)

// BaseRouter
// the very base class to inherit
type BaseRouter struct{}

func (b *BaseRouter) Before(r iface.IRequest) {}

func (b *BaseRouter) Handle(r iface.IRequest) {}

func (b *BaseRouter) After(r iface.IRequest) {}
