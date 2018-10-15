package main

import (
	"config"
	"spider"
	"flag"
	"logger"
	"runtime"
	"github.com/PuerkitoBio/goquery"
	"os"
	"engine"
	"fmt"
	"strconv"
	"visualchina/parser/editorial"
)

var (
	urlInfo config.UrlStruct
	url string
	ok bool
	flagType int
	flagAll  int
)

func init()  {
	//获取命令行参数
	 flag.IntVar(&flagType,"type", 0, "0:编辑图片,1:创意壁纸,2:创意图片,3:设计素材")
	 flag.IntVar(&flagAll,"all",0,"1:抓取所有,0:单个抓取")
	 flag.Parse()
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	urls := config.InitUrls()
	switch flagAll {
	case 1:
		spider.InitAllJobs(urls)
	case 0:
		if urlInfo,ok = urls[flagType];!ok{
              logger.Error.Fatalln("flagType 值非法，没找到相应的URL")
		}
		argReq := engine.RequestArgs{
			Type: strconv.Itoa(flagType),
		}
		spider.InitJobs(urlInfo,argReq)
	}
}


func testgoquery() {
	open, e := os.Open("src/test/editorial.html")
	if e != nil{
		panic(e)
	}

	document, i := goquery.NewDocumentFromReader(open)
	if i != nil {
		panic(i)
	}
	ret := engine.ParseResult{}
	document.Find(".classify-list>li").Each(func(i int, selection *goquery.Selection) {
        a := selection.Find("a")
        title,fbool := a.Attr("title")
        if fbool == true {
			 url,_ := a.Attr("href")
			 test := engine.Request{
				 Url:url,
				 Parser:engine.NewFuncParser( editorial.ParseEditorial,title),
			 }
			 ret.Requests = append(ret.Requests,test)
		}
	})

	 fmt.Printf("%+v",ret.Requests)
}












