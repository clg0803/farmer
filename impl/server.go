package impl

import (
	"farmer/iface"
	"fmt"
	"net"
	"time"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func (s *Server) Start() {
	fmt.Printf(":[START]: Server listen at: %s:%d\n", s.IP, s.Port)
	go s.listenAndServe()
}

func (s *Server) Stop() {
	fmt.Println(":[STOP]: Server: ", s.Name)
}

func (s *Server) Serve() {
	s.Start()
	// TODO: serve()
	for {
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) listenAndServe() {
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
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Print(":[ERR]: ACCEPT FAILED ", err)
			continue
		}
		// TODO:
		go s.work(conn)
	}
}

func (s *Server) work(conn *net.TCPConn) {
	for {
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println(":[ERR]: READ ERR", err)
			continue
		}
		if _, err := conn.Write(buf[:cnt]); err != nil {
			fmt.Println(":[ERR]: ERR WHEN WRITE BACK TO CLIENT", err)
			continue
		}
	}
}

func NewServer(name string) iface.IServer {
	return &Server{Name: name, IPVersion: "tcp4", IP: "0.0.0.0", Port: 8848}
}
