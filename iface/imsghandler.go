package iface

// IMsgHandler
// maintains id->Router mapping and invoke Router by MsgID
type IMsgHandler interface {
	AddRouter(msgId uint32, router IRouter) // ADD to Map
	HandleRequest(r IRequest)               // get handled, identified by r.MsgID()
	AddWorkers()                            // ADD `workerPoolSize` workers
	SendMsgToTaskQueue(r IRequest)          // send msg to task queue to be done by worker
}
