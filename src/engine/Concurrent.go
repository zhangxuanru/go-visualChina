package engine

import (
	"logger"
	"sync"
	"fmt"
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
	  out := make(chan ParseResult,len(seeds))
	  c.Scheduler.Run()
	  ws.Add(len(seeds))
	  for i:=0; i<c.WorkerCount;i++{
          c.CreateWorker(c.Scheduler.WorkerChan(),out)
	  }
	  for _,r := range seeds{
	  	  c.Scheduler.Submit(r)
	  }
	  c.WorkerRun(out)
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
		   out <- result
		}
	}()
}

func (c *ConcurrentEngine) WorkerRun(out chan ParseResult) {
	go func() {
		for{
			ret := <-out
			for _, request := range ret.Requests {
				result,err := Fetch(request)
				if err!=nil{
					logger.Error.Panic("抓取",request.Url,"失败，失败原因是:",err)
				}
				fmt.Println(result)
			}
			ws.Done()
		}
	}()
}



