package scheduler

import (
	"engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan  chan engine.Request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	 return q.requestChan
}

func (q *QueuedScheduler) Run()  {
	q.requestChan = make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(r engine.Request)  {
	go func() {
		q.requestChan <- r
	}()
}

