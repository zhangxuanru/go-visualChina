package parser

import (
	"engine"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"logger"
	"visualchina/persist"
	"fmt"
	"visualchina/Model"
)

var editorial = persist.Editorial{}

/*
抓取 编辑图片首页内容
*/
func ParseEditorial(contents []byte, url string,args engine.RequestArgs) engine.ParseResult  {
	reader := strings.NewReader(string(contents))
	document, i := goquery.NewDocumentFromReader(reader)
	if i != nil {
		logger.Error.Println("grab url ",url," args:",args," goquery error:",i)
	}
	ret := engine.ParseResult{}
	document.Find(".classify-list>li").Each(func(i int, selection *goquery.Selection) {
		a := selection.Find("a")
		title,f := a.Attr("title")
		if f == true {
			args.Title = title
			url,_ := a.Attr("href")
			req := engine.Request{
				Url:url,
				Parser:engine.NewFuncParser(ParseEditorialLeftNav,title),
				Args:args,
			}
			ret.Requests = append(ret.Requests,req)
		}
	})
	return ret
}


/*
处理编辑图片 左侧栏目数据
*/
func ParseEditorialLeftNav(contents []byte,url string,args engine.RequestArgs) engine.ParseResult {

	//test 如果要拿到返回数据，就不能用go routine
	ret := saveTestLeftNav(url,args)
	fmt.Println("ret:",ret)
	return engine.ParseResult{}
	//test

	bo := make(chan bool)
	fmt.Println(url)
    go saveLeftNav(url,args,bo)
	go editorial.NavDbDataRun()
	<-bo
	 return engine.ParseResult{}
}

//保存左侧栏目数据
func saveLeftNav(url string,args engine.RequestArgs, b chan bool) {
	editorial.NavRun(10)
	editorial.NavDbRun()
	editorial.NavSubmit(persist.NavStruct{
		Title:args.Title,
		Type:args.Type,
		Url:url,
	})
	save := editorial.NavSave()
	b <- true
	editorial.NavDbSubmit(save)
}


func saveTestLeftNav(url string,args engine.RequestArgs) Model.NavDb{
	nav:= persist.NavStruct{
		Title: args.Title,
		Type:  args.Type,
		Url:   url,
	}
	save := editorial.NavTestSave(nav)
    return save
}



