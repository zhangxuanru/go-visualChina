package editorial

import (
	"engine"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"logger"
	 "visualchina/persist/editorialPersist"
	"visualchina/Model"
	"strconv"
	"time"
	"libary/query"
	"libary/upload"
	"fmt"
)

var editorial = editorialPersist.Editorial{}

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
	var  parser  engine.ParserFunc

	navList := query.ParseEditorial(document)
	for _,nav := range navList{
		url := nav.Url
		args.Title = nav.Title
		result := saveNav(url,args)
		parser = ParseEditorialNavLevelPage
		//原创
		if strings.Contains(url,"original"){
			parser = ParseEditorialNavOriginal
		}
		//滚动
		if strings.Contains(url,"editorial-update"){
			parser = ParseEditorialUpdateOriginal
		}
		//专题
		if strings.Contains(url,"topics"){
			parser = ParseEditorialNavTopic
		}
		req := engine.Request{
			Url:url,
			Parser:engine.NewFuncParser(parser,nav.Title),
			Args: engine.RequestArgs{
				Id:result.Id,
				Type:args.Type,
				Update:args.Update,
				Title:result.Title,
			},
		}
		ret.Requests = append(ret.Requests,req)
	}
	//获取首页banner

	//获取首页推荐专题

	return ret
}


//滚动二级页（将分类ID保存到导航表中）
func ParseEditorialUpdateOriginal(contents []byte,url string,args engine.RequestArgs) engine.ParseResult {
	if args.Update == "1"{ //如果只抓最新的数据则不需要做以下的操作
		 return engine.ParseResult{}
	}
	fmt.Println("抓取",args.Title,"---栏目开始:")
	reader := strings.NewReader(string(contents))
	document, e := goquery.NewDocumentFromReader(reader)
	if e != nil {
		logger.Error.Println("grab url ",url," args:",args," goquery error:",e)
	}
	catList := query.ParseEditorialUpdate(document)
	for catId,title := range catList{
		nav:= Model.NavDb{
			Title: title,
		    CategoryId:catId,
		}
		nav.UpdateNavCategoryIdByTitle()
	}
	return engine.ParseResult{}
}


//抓取栏目 二级页
func ParseEditorialNavLevelPage(contents []byte,url string,args engine.RequestArgs) engine.ParseResult {
	if args.Update == "1"{ //如果只抓最新的数据则不需要做以下的操作
		return getGroupDataByNavId(args.Id)
	}
	reader := strings.NewReader(string(contents))
	document, e := goquery.NewDocumentFromReader(reader)
	if e != nil {
		logger.Error.Println("grab url ",url," args:",args," goquery error:",e)
	}
	tagList := query.ParseEditorialLevelPage(document)
	//保存tag
	for _,tag := range tagList{
		  tag.Type = args.Type
          editorial.SaveTag(tag)
	}
	//保存栏目页面底部推广数据
	generalizeList := query.ParseEditorialLevelGeneralize(document)
	for _,genera := range generalizeList{
		imageId := upload.UploadToQiniu(genera.Src)
		genera.ImageId = strconv.FormatInt(imageId,10)
		editorial.SaveGenera(genera)
	}
	//保存栏目页上面的推荐数据
	levelRecommend := query.ParseEditorialLevelRecommend(document)
	for _,recommend := range levelRecommend{
           editorial.SaveRecommend(recommend)
	}
	return  getGroupDataByNavId(args.Id)
}


//保存左侧栏目数据
func saveNav(url string,args engine.RequestArgs) Model.NavDb{
	iType, _ := strconv.Atoi(args.Type)
	nav:= Model.NavDb{
		Title: args.Title,
		Type:  iType,
		GrabUrl:   url,
		ExecDate: time.Now().Unix(),
	}
	save := editorial.NavSave(nav)
	return save
}

