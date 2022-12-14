package impl

import (
	"farmer/iface"
	"fmt"
	"net"
	"time"
)

type Server struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int
	msgHandler  iface.IMsgHandler
	connManager ConnManager

	MaxConn     int32
	MaxPackSize int32
}

func (s *Server) Start() {
	fmt.Printf(":[START]: Server listen at: %s:%d\n", s.IP, s.Port)
	go s.listenAndServe()
}

func (s *Server) Stop() {
	fmt.Println(":[STOP]: Server: ", s.Name)
	s.connManager.CleanAllConn()
}

func (s *Server) Serve() {
	s.Start()
	// TODO: serve()
	// do sth. while handling requests
	for {
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) AddRouter(msgId uint32, r iface.IRouter) {
	s.msgHandler.AddRouter(msgId, r)
	fmt.Println(":[SUCCESS]: ADD A NEW ROUTER!")
}

func (s *Server) GetConnMgr() iface.IConnectManager {
	return &s.connManager
}

func (s *Server) listenAndServe() {
	s.msgHandler.AddWorkers()
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println(":[ERR]: RESOLVE TCP ADDR ", err)
		return
	}
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println(":[ERR]: CAN'T ESTABLISH LISTENER")
		return
	}
	fmt.Println(":[START]: FARMER <", s.Name, "> BEGIN TO LABOR, NOW LISTENING ... ")
	fmt.Printf(":[CONF]: IPVERSION:%s ADDR:%s PORT:%d MAX_CONN:%d MAX_PACKSIZE:%d \n",
		s.IPVersion, s.IP, s.Port, s.MaxConn, s.MaxPackSize,
	)

	var connID uint32 = 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Print(":[ERR]: ACCEPT FAILED ", err)
			continue
		}
		if s.connManager.ConnectedNum() >= s.MaxConn {
			conn.Close()
			continue
		}
		// create a 'connection' obj and bind task to it
		estConn := NewConnection(s, conn, connID, s.msgHandler) // assume no bug in NewConnection()
		go estConn.Start()
		connID++
	}
}

func NewServer(name string) iface.IServer {
	return &Server{
		Name:        name,
		IPVersion:   "tcp4",
		IP:          "127.0.0.1",
		Port:        8848,
		msgHandler:  NewMsgHandler(5, 1024),
		connManager: *NewConnManager(),

		MaxConn:     1024,
		MaxPackSize: 65535,
	}
}
