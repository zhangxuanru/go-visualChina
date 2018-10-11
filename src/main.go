package main

import (
	"config"
	"spider"
	"fmt"
)

var (
	urlInfo config.UrlStruct
	url string
	ok bool
)

func main(){
	urls := config.InitUrls()
	switch config.FlagAll {
	case 1:
		spider.InitAllJobs(urls)
	case 0:
		if urlInfo,ok = urls[config.FlagType];!ok{
           config.Error.Fatalln("flagType 值非法，没找到相应的URL")
		}
		spider.InitJobs(urlInfo)
	}
	fmt.Println(url)
}













