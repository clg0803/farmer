package impl

import (
	"errors"
	"farmer/iface"
	"fmt"
	"sync"
)

type ConnManager struct {
	connections map[uint32]iface.IConnection
	cl          sync.RWMutex
}

func (cm *ConnManager) Add(conn iface.IConnection) {
	cm.cl.Lock()
	defer cm.cl.Unlock()
	cm.connections[conn.GetConnectionID()] = conn
	fmt.Println(":[SUCCESS]: CONN ADDED TO CONN_MANAGER, CONNECTED NUM = ", cm.ConnectedNum())
}

func (cm *ConnManager) Remove(conn iface.IConnection) {
	cm.cl.Lock()
	defer cm.cl.Unlock()
	delete(cm.connections, conn.GetConnectionID())
	fmt.Println(":[SUCCESS]: CONNECTION ", conn.GetConnectionID(), " REMOVED", ", CONNECTED NUM = ", cm.ConnectedNum())
}

func (cm *ConnManager) Get(connID uint32) (iface.IConnection, error) {
	cm.cl.Lock()
	defer cm.cl.Unlock()
	if conn, ok := cm.connections[connID]; !ok {
		return nil, errors.New("CONNECTION NOT FOUND")
	} else {
		return conn, nil
	}
}

func (cm *ConnManager) ConnectedNum() int32 {
	return int32(len(cm.connections))
}

func (cm *ConnManager) CleanAllConn() {
	cm.cl.Lock()
	defer cm.cl.Unlock()
	for connID, conn := range cm.connections {
		conn.Stop()
		delete(cm.connections, connID)
	}
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]iface.IConnection),
	}
}
