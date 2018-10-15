package spider

import (
	"engine"
	"scheduler"
	"config"
	"logger"
)

//抓取所有页面
func InitAllJobs(urls map[int]config.UrlStruct)  {
	
}

//抓取单页面
func InitJobs(urlinfo config.UrlStruct,argReq engine.RequestArgs)  {
	logger.Info.Println("start grab ",urlinfo.Url)
	e := engine.ConcurrentEngine{
		WorkerCount: 50,
		Scheduler : &scheduler.QueuedScheduler{},
	}
	e.Run(
      engine.Request{
		   Url:    urlinfo.Url,
		   Parser: engine.NewFuncParser(urlinfo.ParseFunc,urlinfo.Name),
		   Args: argReq,
	  })
}
