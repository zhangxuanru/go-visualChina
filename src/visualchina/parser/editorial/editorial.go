package editorial

import (
	"engine"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"logger"
	"visualchina/persist"
	"visualchina/Model"
	"fmt"
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
 	ret := saveLeftNav(url,args)
 	result := engine.ParseResult{}
 	if ret.Id == 0 {
		return result
	}
	req := engine.Request{
           Url:ret.Url,
		   Args:engine.RequestArgs{
               Id:ret.Id,
			   Data:ret,
			},
			Parser:engine.NewFuncParser(ParseEditorialLeftNavData,ret.Title),
	}
	 result.Items = append(result.Items,engine.Item{
		Data:ret,
 	})
	result.Requests = append(result.Requests,req)
	return result
}


func ParseEditorialLeftNavData(contents []byte,url string,args engine.RequestArgs) engine.ParseResult {
	fmt.Println("ParseEditorialLeftNavData:",url,"args",args)
	return engine.ParseResult{}
}




//保存左侧栏目数据
func saveLeftNav(url string,args engine.RequestArgs) Model.NavDb{
	nav:= persist.NavStruct{
		Title: args.Title,
		Type:  args.Type,
		Url:   url,
	}
	save := editorial.NavSave(nav)
    return save
}



