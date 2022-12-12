package iface

// IRouter
// OPEN for customizing, impl 3 method in `IRouter` will finally affect func called by `connection`
type IRouter interface {
	Before(request IRequest)
	Handle(request IRequest)
	After(request IRequest)
}
