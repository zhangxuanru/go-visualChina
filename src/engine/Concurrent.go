package engine

import (
	"fetcher"
	"config"
	"fmt"
)

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
      for i:=0; i < c.WorkerCount;i++{
      	   c.CreateWorker(c.Scheduler.WorkerChan(),out)
	  }
	 for _,r := range seeds{
    	 c.Scheduler.Submit(r)
	 }
	for {
          result := <-out
          fmt.Println(result)
	}
}


func (c *ConcurrentEngine) CreateWorker(in chan Request, out chan ParseResult)  {
	go func() {
       for{
           request := <-in
		   result,err := fetcher.Fetch(request)
		   if err!=nil{
		    	config.Error.Panic("抓取",request.Url,"失败，失败原因是:",err)
		   }
		   out <- result
	   }
	}()
}



