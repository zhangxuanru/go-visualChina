package parser

import (
	"engine"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"logger"
	"visualchina/persist"
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
	argReq := engine.RequestArgs{}
	document.Find(".classify-list>li").Each(func(i int, selection *goquery.Selection) {
		a := selection.Find("a")
		title,f := a.Attr("title")
		if f == true {
			argReq.Title = title
			url,_ := a.Attr("href")
			req := engine.Request{
				Url:url,
				Parser:engine.NewFuncParser(ParseEditorialLeftNav,title),
				Args:argReq,
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
	go saveLeftNav(url,args)

	return engine.ParseResult{}
}

//保存左侧栏目数据
func saveLeftNav(url string,args engine.RequestArgs)  {
	editorial.NavRun(10)
	editorial.NavSubmit(persist.NavStruct{
		Title:args.Title,
		Url:url,
	})
	editorial.NavSave()
}




