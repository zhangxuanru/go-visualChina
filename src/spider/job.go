package spider

import (
	"engine"
	"scheduler"
	"config"
)

//抓取所有页面
func InitAllJobs(urls map[int]config.UrlStruct)  {
	
}

//抓取单页面
func InitJobs(urlinfo config.UrlStruct)  {
	e := engine.ConcurrentEngine{
		WorkerCount: 10,
		Scheduler : &scheduler.QueuedScheduler{},
	}
	e.Run(
      engine.Request{
		   Url:    urlinfo.Url,
		   Parser: engine.NewFuncParser(urlinfo.ParseFunc,urlinfo.Name),
	  })
}
