package editorial

import (
	"engine"
	"visualchina/persist/editorialPersist"
	"strconv"
	"config/constant"
)

//专题二级页
func ParseEditorialNavTopic(contents []byte,url string,args engine.RequestArgs) (ret engine.ParseResult) {
	if len(contents) == 0{
		return  engine.ParseResult{}
	}
	topicData, err := editorialPersist.ParseTopicJson(contents)
	if err != nil {
		return engine.ParseResult{}
	}
    editorial.SaveTopic(topicData)
    if editorial.TopicSaveList == nil{
   	     return engine.ParseResult{}
     }
     for itemId,val := range editorial.TopicSaveList{
		 id := strconv.FormatInt(itemId, 10)
		 if val.Type == 1{ //分类
			 ret.Requests = append(ret.Requests,engine.Request{
				 Url: constant.BaseUrl+"/editorial-topics-category?per_page=100&page=1&cid="+id,
				 Method:"GET",
				 Parser:engine.NewFuncParser(ParseEditorialTopicCategory,args.Title),
				 Args:engine.RequestArgs{
					 CategoryId:itemId,
					 Page:int64(1),
					 UrlType:"topiccategory",
				 },
			 })
		 }
		 if val.Type == 0{ //topic
			 ret.Requests = append(ret.Requests,engine.Request{
				 Url: constant.BaseUrl+"/topic/"+id,
				 Method:"GET",
				 Parser:engine.NewFuncParser(ParseEditorialTopicGroup,args.Title),
				 Args:engine.RequestArgs{
					 CategoryId:itemId,
					 Page:1,
					 UrlType:"topicgroup",
				 },
			 })
		 }
	}
  editorial.TopicSaveList = nil
  return ret
}

//解析专题分类下的各个专题
func ParseEditorialTopicCategory(contents []byte,url string,args engine.RequestArgs) (ret engine.ParseResult) {
	if len(contents) == 0{
		return  engine.ParseResult{}
	}
	topicData, err := editorialPersist.ParseTopicCategoryJson(contents)
	if err != nil {
		 return engine.ParseResult{}
	}
	if len(topicData.Data.Data.List) == 0{
		return engine.ParseResult{}
	}
	for _,val := range topicData.Data.Data.List {
		_,err := editorial.SaveCategoryItem(val)
		if err == nil{
			ret.Requests = append(ret.Requests,engine.Request{
				Url: constant.BaseUrl+"/topic/"+ strconv.Itoa(val.ID),
				Method:"GET",
				Parser:engine.NewFuncParser(ParseEditorialTopicGroup,args.Title),
				Args:engine.RequestArgs{
					CategoryId:int64(val.ID),
					Page:1,
					UrlType:"topicgroup",
				},
			})
		}
	}
	if (args.Page)*100 >= topicData.Data.Data.TotalCount{
		 return engine.ParseResult{}
	}
	if (args.Page)*100 < topicData.Data.Data.TotalCount{
		categoryId :=  strconv.FormatInt(args.CategoryId,10)
		nextPage :=  strconv.FormatInt(args.Page+1,10)
		ret.Requests = append(ret.Requests,engine.Request{
			Url: constant.BaseUrl+"/editorial-topics-category?per_page=100&page="+nextPage+"&cid="+categoryId,
			Method:"GET",
			Parser:engine.NewFuncParser(ParseEditorialTopicCategory,args.Title),
			Args:engine.RequestArgs{
				CategoryId:args.CategoryId,
				Page:args.Page+1,
				UrlType:"topiccategory",
			},
		})
	}
	return ret
}

//解析专题下的group
func ParseEditorialTopicGroup(contents []byte,url string,args engine.RequestArgs) (ret engine.ParseResult) {
	if len(contents) == 0{
		return  engine.ParseResult{}
	}
	topic, err := editorialPersist.ParseTopicGroupJson(contents)
	if err != nil{
		return engine.ParseResult{}
	}
	TopicImages := topic.TopicInfo.TopicImages
	for _,item := range TopicImages.List{
		item.Cid = args.CategoryId
	    editorial.SaveTopicGroup(item)
	}

	if TopicImages.PageNum * TopicImages.PageSize >= TopicImages.Total{
		return engine.ParseResult{}
	}
	 categoryId := strconv.FormatInt(args.CategoryId,10)
	 nextPage := strconv.FormatInt(args.Page +1,10)
	 ret.Requests = append(ret.Requests,engine.Request{
			Url: constant.BaseUrl+"/topic/"+categoryId+"?showType=1&pageNum="+nextPage+"&pageSize=100",
			Method:"GET",
			Parser:engine.NewFuncParser(ParseEditorialTopicGroup,args.Title),
			Args:engine.RequestArgs{
				CategoryId:args.CategoryId,
				Page:args.Page+1,
				UrlType:"topicgroup",
			},
		})
	return ret
}




