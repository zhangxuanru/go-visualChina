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
	//testeditorupdate()
	//return

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


//编辑图片首页匹配
func testgoquery() {
	open, e := os.Open("src/test/editorial.html")
	if e != nil{
		panic(e)
	}

	document, i := goquery.NewDocumentFromReader(open)
	if i != nil {
		panic(i)
	}
	//ret := engine.ParseResult{}
	document.Find(".classify-list>li").Each(func(i int, selection *goquery.Selection) {
        a := selection.Find("a")
        title,_ := a.Attr("title")
		url,_ := a.Attr("href")
		fmt.Println(title,url)
		//if fbool == true {
		//	 url,_ := a.Attr("href")
		//	 test := engine.Request{
		//		 Url:url,
		//		 Parser:engine.NewFuncParser( editorial.ParseEditorial,title),
		//	 }
		//	 ret.Requests = append(ret.Requests,test)
		//}
	})

	// fmt.Printf("%+v",ret.Requests)
}

//编辑图片--栏目页
func testeditorupdate()  {
	file, e := os.Open("src/test/editorial_update.html")
	if e != nil{
		panic(e)
	}
	document, i := goquery.NewDocumentFromReader(file)
	if i != nil {
		panic(i)
	}
	document.Find(".indexnav-tabs>li").Each(func(i int, selection *goquery.Selection) {
		a := selection.Find("a")
		title,_ := a.Attr("title")
		url,_ := a.Attr("href")
		fmt.Println(title,"https://www.vcg.com"+url)
	})
}










