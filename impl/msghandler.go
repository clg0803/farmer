package impl

import (
	"farmer/iface"
	"fmt"
	"strconv"
)

type MsgHandler struct {
	routerTable    map[uint32]iface.IRouter
	WorkerPoolSize uint32
	taskQueue      []chan iface.IRequest // one worker has one taskQueue
	TaskQueueSize  uint32
}

func (mh *MsgHandler) AddRouter(msgId uint32, router iface.IRouter) {
	if _, ok := mh.routerTable[msgId]; ok {
		panic(":[FETAL]: REPEATED ROUTER, MSG_ID = " + strconv.Itoa(int(msgId)))
	}
	mh.routerTable[msgId] = router
	fmt.Println(":[SUCCESS]: ADD ROUTER, MSG_ID = ", msgId)
}

func (mh *MsgHandler) HandleRequest(r iface.IRequest) {
	handler, ok := mh.routerTable[r.GetMsgID()]
	if !ok {
		fmt.Println(":[ERR]: ROUTER NOT FOUND FOR MSG_ID = ", r.GetMsgID())
		return
	}
	handler.Before(r)
	handler.Handle(r)
	handler.After(r)
}

func (mh *MsgHandler) AddWorkers() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// one worker is mapped to one task queue
		mh.taskQueue[i] = make(chan iface.IRequest, mh.TaskQueueSize)
		go mh.startAWorker(i, &mh.taskQueue[i])
	}
}

func (mh *MsgHandler) SendMsgToTaskQueue(r iface.IRequest) {
	workerID := r.GetConnection().GetConnectionID() % mh.WorkerPoolSize
	mh.taskQueue[workerID] <- r
	fmt.Println(":[SUCCESS]: ADD TO A WORKER, ID = ", workerID,
		"CONNECTION ID = ", r.GetConnection().GetConnectionID(),
		"MESSAGE ID = ", r.GetMsgID())
}

func (mh *MsgHandler) startAWorker(workerId int, tq *chan iface.IRequest) {
	fmt.Println(":[SUCCESS]: WORKER START, ID = ", workerId)
	for {
		select {
		case req := <-*tq:
			mh.HandleRequest(req)
		}
	}
}

func NewMsgHandler(ps, ts uint32) *MsgHandler {
	return &MsgHandler{
		routerTable:    make(map[uint32]iface.IRouter),
		WorkerPoolSize: ps,
		taskQueue:      make([]chan iface.IRequest, ps),
		TaskQueueSize:  ts,
	}
}
