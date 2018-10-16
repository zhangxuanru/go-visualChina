package editorial

import (
	"engine"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"logger"
	"visualchina/persist"
	"visualchina/Model"
	"strconv"
	"fmt"
)

var editorial = persist.Editorial{}

const BaseUrl  = "https://www.vcg.com"

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
	//首页左边导航处理 start
	document.Find(".classify-list>li").Each(func(i int, selection *goquery.Selection) {
		a := selection.Find("a")
		title,f := a.Attr("title")
		if f == true {
			args.Title = title
			url,_ := a.Attr("href")
			result := saveLeftNav(url,args)
			req := engine.Request{
				Url:url,
				Parser:engine.NewFuncParser(ParseEditorialNavLevelPage,title),
				Args: engine.RequestArgs{
				       Id:result.Id,
				       Type:args.Type,
				       Data:result,
			    },
			}
			ret.Items = append(ret.Items,engine.Item{
				Data:result,
			})
			ret.Requests = append(ret.Requests,req)
		}
	})
	//首页左边导航处理 end
	return ret
}

//抓取栏目 二级页
func ParseEditorialNavLevelPage(contents []byte,url string,args engine.RequestArgs) engine.ParseResult {
	reader := strings.NewReader(string(contents))
	document, e := goquery.NewDocumentFromReader(reader)
	if e != nil {
		logger.Error.Println("grab url ",url," args:",args," goquery error:",e)
	}
	ret := engine.ParseResult{}
	//保存子导航栏目并设置子导航栏目函数start
	document.Find(".indexnav-tabs>li").Each(func(i int, selection *goquery.Selection) {
		a := selection.Find("a")
		title,f := a.Attr("title")
		if f == true {
			args.Title = title
			args.Pid   = args.Id
			url,_ := a.Attr("href")
			url = BaseUrl+url
			result := saveLeftNav(url,args)
			req := engine.Request{
				Url:url,
				Parser:engine.NewFuncParser(ParseEditorialPageNavData,title),
				Args:engine.RequestArgs{
					Id: result.Id,
					Pid: args.Id,
					Type:args.Type,
				},
			}
			ret.Requests = append(ret.Requests,req)
			fmt.Println()
			fmt.Printf("%s#%s#%+v",title,url,req.Args)
		}
	})
	//保存子导航栏目并设置子导航栏目函数end
	return engine.ParseResult{}
}


//处理二级子导航页面
func ParseEditorialPageNavData(contents []byte,url string,args engine.RequestArgs) engine.ParseResult {
	//抓取图集start

	//抓取图集end
	return engine.ParseResult{}
}


//保存左侧栏目数据
func saveLeftNav(url string,args engine.RequestArgs) Model.NavDb{
	iType, _ := strconv.Atoi(args.Type)
	nav:= Model.NavDb{
		Title: args.Title,
		Type:  iType,
		Url:   url,
		Pid:   args.Pid,
	}
	save := editorial.NavSave(nav)
	return save
}



