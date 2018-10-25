package engine

import (
	"logger"
	"sync"
)

var ws sync.WaitGroup
var group sync.WaitGroup

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
	  count := len(seeds)
	  out := make(chan ParseResult)
	  c.Scheduler.Run()
	  for i:=0; i<c.WorkerCount;i++{
          c.CreateWorker(c.Scheduler.WorkerChan(),out,&count)
	  }
	  for _,r := range seeds{
	  	   c.Scheduler.Submit(r)
	  }
	 for{
		ret := <-out
		for _, request := range ret.Requests {
			c.Scheduler.Submit(request)
			count++
		}
		if count == 0{
			break
		}
	}
}


func (c *ConcurrentEngine) CreateWorker(in chan Request, out chan ParseResult,req *int)  {
	go func() {
		for{
           request := <-in
		   result,err := FetchUrl(request)
		   if err!=nil{
			    logger.Error.Println("抓取",request.Url,"失败，失败原因是:",err)
			    continue
		   }
		    out <- result
			*req--
		}
	}()
}

