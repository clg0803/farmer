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
	var connID uint32 = 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Print(":[ERR]: ACCEPT FAILED ", err)
			continue
		}
		// create a 'connection' obj and bind task to it
		estConn := NewConnection(conn, connID, wb) // assume no bug in NewConnection()
		go estConn.Start()
		connID++
	}
}

// bind each connection with this function
// will customize it(say in CLIENT) in future
func wb(conn *net.TCPConn, data []byte, length int) error {
	if _, err := conn.Write(data[:length]); err != nil {
		fmt.Println(":[ERR]: ERR WHEN WRITE BACK TO CLIENT", err)
		return err
	}
	fmt.Println(":[SUCCESS]: CALLBACK FUNC FINISHED")
	return nil
}

func NewServer(name string) iface.IServer {
	return &Server{Name: name, IPVersion: "tcp4", IP: "0.0.0.0", Port: 8848}
}
