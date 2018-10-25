package query

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"net/url"
	"strconv"
	"config/constant"
)

type NavstrUct struct {
	Title string
	Url   string
}

type TagModel struct {
    Id string
    Type string
    Pid string
    Code string
    Name string
    Url string
    Index string
}

type TagModels struct {
	TagModel
	SubTags []TagModel
}

type Generalize struct {
	CategoryId  string
	ImageId     string
 	TopicId     string
 	Gtype       string
	Title       string
	Src         string
	Link        string
}

type LevelRecommend struct {
        GroupId  string
        TopicId  string
	    CategoryId  string
}

//提取编辑图片栏目首页左边导航处理
func ParseEditorial(document *goquery.Document) (mp map[int]NavstrUct) {
	mp = make(map[int]NavstrUct)
	k:= 0
	document.Find(".classify-list>li").Each(func(i int, selection *goquery.Selection) {
		a := selection.Find("a")
		title,f := a.Attr("title")
		if f == true {
			url,_ := a.Attr("href")
			nav := NavstrUct{
				Title:title,
				Url:url,
			}
		    mp[k] = nav
		    k++
		}
	})
	return mp
}


//提取栏目页获取tag标签
func ParseEditorialLevelPage(document *goquery.Document) (mp map[int]TagModels) {
	mp = make(map[int]TagModels)
	k:=0
	document.Find(".channelv-nav-tag>.channelNav").Each(func(i int, selection *goquery.Selection) {
		id,_:= selection.Attr("data-id")
		pid,_ := selection.Attr("data-pid")
		code,_ := selection.Attr("data-code")
		name,_ := selection.Attr("data-name")
		tags := TagModels{}
		tags.Id = id
		tags.Pid = pid
		tags.Name = name
		tags.Code = code
		if i >0 {
		  document.Find(".channel-open>.channelNav-ul-1").Eq(i-1).Find("li").Each(func(j int, selection *goquery.Selection) {
				subId,_ := selection.Attr("data-id")
				subPid,_ := selection.Attr("data-pid")
				subName,_ := selection.Attr("data-name")
				subIndex,_ := selection.Attr("data-index")
			    subTag:= TagModel{}
			    subTag.Id = subId
			    subTag.Pid=subPid
			    subTag.Name=subName
			    subTag.Index = subIndex
				tags.SubTags = append(tags.SubTags,subTag)
             })
		}
		mp[k] = tags
		k++
	})
	return mp
}


//获取栏目页面下面的推广专题
func ParseEditorialLevelGeneralize(document *goquery.Document) (mp map[int]Generalize) {
	categoryId, _:= document.Find(".channelv-nav-tag>.channelNav").Eq(0).Attr("data-id")
	mp = make(map[int]Generalize)
	k := 0
	document.Find(".gen-item").Each(func(i int, selection *goquery.Selection) {
          link := selection.Find("a")
		  href,_ := link.Attr("href")
		  src,_ := link.Find("img").Attr("src")
		  text:= link.Text()
		  topicId:=""
		  if strings.Contains(href,"topic"){
				topicId = strings.TrimLeft(href, "/topic/")
		  }
		  mp[k] = Generalize{
				CategoryId:categoryId,
			    TopicId:topicId,
			    Gtype:"0",
			    Title:text,
			    Src:src,
			    Link:href,
			}
			k++
	})
	document.Find(".link-wraper>a").Each(func(i int, selection *goquery.Selection) {
		href,_ := selection.Attr("href")
		title:= selection.Text()
		topicId:=""
		if strings.Contains(href,"topic"){
			topicId = strings.TrimLeft(href, "/topic/")
		}
		mp[k] = Generalize{
			CategoryId:categoryId,
			TopicId:topicId,
			Gtype:"1",
			Title:title,
			Link:href,
		}
		k++
	})
	return  mp
}

//获取栏目页面上面的推荐数据
func ParseEditorialLevelRecommend(document *goquery.Document) (mp map[int]LevelRecommend) {
	categoryId, _:= document.Find(".channelv-nav-tag>.channelNav").Eq(0).Attr("data-id")
	mp = make(map[int]LevelRecommend)
	k := 0
	document.Find(".picrecommend>.pro-item>.pro-item-box").Each(func(i int, selection *goquery.Selection) {
		link,_ := selection.Find("a").Attr("href")
		link = strings.TrimSpace(link)
		topicId := ""
		groupId := ""
		if strings.Contains(link,"/topic/"){
			topicId = strings.TrimLeft(link,"/topic/")
		}
		if strings.Contains(link,"/group/"){
			groupId = strings.TrimLeft(link,"/group/")
		}
		mp[k] = LevelRecommend{
          CategoryId:categoryId,
          GroupId:groupId,
          TopicId:topicId,
		}
        k++
	})
	return  mp
}








//获取滚动栏目下的分类列表
func ParseEditorialUpdate(document *goquery.Document)(mp map[int64]string)  {
	mp = make(map[int64]string)
	document.Find(".indexnav-tabs>li>a").Each(func(i int, selection *goquery.Selection) {
		href,_ := selection.Attr("href")
		uri, _ := url.Parse(constant.BaseUrl+href)
		category, _ := url.ParseQuery(uri.RawQuery)
	    categoryId := category.Get("category")
		catId,_ := strconv.ParseInt(categoryId,10,64)
		mp[catId] = selection.Text()
	})
	return mp
}

