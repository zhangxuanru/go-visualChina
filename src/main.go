package main

import (
	"spider"
	"fmt"
)

var (
	url string
	ok bool
)

func main(){
	urls := initUlrs()
	switch flagAll {
	case 1:
		spider.InitAllJobs(urls)
	case 0:
		if url,ok = urls[flagType];!ok{
            Error.Fatalln("flagType 值非法，没找到相应的URL")
		}
		spider.InitJobs(url)
	}
	fmt.Println(url)
}













