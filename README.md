# farmer
a webserver model for extension

# example
run `impl/server_test.go` as *Client*, run `main.go` as *Server*

# how to customize
## Server
1. create a `Server` obj with config difined in `server.go` and `connection.go`, `SetOnConnStart()` and `SetOnConnEnd()` can be use when a connection begin/end
2. follow `main.go` example, define a `Router` that extends Base Router, and add it to `Server` obj through func `AddRouter(int, IRouter)` with `msgID`

## Client
follow `impl/server_test.go`, simply use `net.Dial()` to create a connection and send packed data(by Pack.pack()) to `Server`
