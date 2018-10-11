package engine

import (
	"logger"
	"fmt"
	"sync"
)
var ws sync.WaitGroup

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
}

type Scheduler interface {
	WorkerChan() chan Request
	Run()
	Submit(request Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request)  {
	  out := make(chan ParseResult)
	  c.Scheduler.Run()
	  ws.Add(len(seeds))
	  for i:=0; i<c.WorkerCount;i++{
          c.CreateWorker(c.Scheduler.WorkerChan(),out)
	  }
	  for _,r := range seeds{
	  	  c.Scheduler.Submit(r)
	  }
	  ws.Wait()
}


func (c *ConcurrentEngine) CreateWorker(in chan Request, out chan ParseResult)  {
	go func() {
		for{
           request := <-in
		   result,err := Fetch(request)
		   if err!=nil{
			    logger.Error.Panic("抓取",request.Url,"失败，失败原因是:",err)
		   }
		  // out <- result

		   fmt.Println("rrr:",result)
		   fmt.Println(request)
		   ws.Done()
		}
	}()
}



