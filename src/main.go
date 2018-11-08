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
	"io/ioutil"
	"errors"
	"strings"
	"regexp"
)

var (
	urlInfo config.UrlStruct
	url string
	ok bool
	flagType int
	flagAll  int
	flagUpdate int
)

func init()  {
	//获取命令行参数
	 flag.IntVar(&flagType,"type", 0, "0:编辑图片,1:创意壁纸,2:创意图片,3:设计素材")
	 flag.IntVar(&flagAll,"all",0,"1:抓取所有,0:单个抓取")
	 flag.IntVar(&flagUpdate,"update",0,"1:只抓取最新数据,0:全部抓取")
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
			Update:strconv.Itoa(flagUpdate),
		}
		spider.InitJobs(urlInfo,argReq)
	}
	fmt.Println("RUN OK")
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
	document.Find(".classify-list>li").Each(func(i int, selection *goquery.Selection) {
        a := selection.Find("a")
        title,_ := a.Attr("title")
		url,_ := a.Attr("href")
		fmt.Println(title,url)
	})
}


//编辑图片--栏目页
func testeditorupdate()  {
	file, e := os.Open("src/test/group.html")
	if e != nil{
		panic(e)
	}
	defer file.Close()
	compile, err := regexp.Compile(`<script>window.__REDUX_STATE__(.*)</script>`)
	if err != nil{
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(file)
	find := compile.FindSubmatch(bytes)
	if len(find)<2{
		panic(errors.New("error"))
	}
    content := strings.TrimSpace(string(find[1]))
	content = strings.TrimFunc(content, func(r rune) bool {
		return r ==' ' || r ==';' || r == '='
	})

    fmt.Println(content)

	//document, _ := goquery.NewDocumentFromReader(file)
	//page := document.Find(".page").Eq(0).Text()
	//page = strings.Trim(page," / ")
	//fmt.Println("---------------")
	//fmt.Println( page)
	//fmt.Println("---------------")



}










