package scheduler

import (
	"engine"
	"fmt"
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
	fmt.Println("r:",r)
       q.requestChan <- r
}

