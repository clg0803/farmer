package impl

import (
	"farmer/iface"
	"fmt"
)

type MsgHandler struct {
	routerTable map[uint32]iface.IRouter
}

func (mh *MsgHandler) AddRouter(msgId uint32, router iface.IRouter) {
	if _, ok := mh.routerTable[msgId]; ok {
		panic(":[FETAL]: REPEATED ROUTER, MSG_ID = " + string(int(msgId)))
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

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		routerTable: make(map[uint32]iface.IRouter),
	}
}
