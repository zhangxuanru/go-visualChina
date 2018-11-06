package editorial

import (
	"engine"
	"visualchina/persist/editorialPersist"
	"config/constant"
	"strconv"
)

//图集页
func ParseEditorialAtlasTopic(contents []byte,url string,args engine.RequestArgs) (ret engine.ParseResult) {
	 atlas, err := editorialPersist.ParseAtlasJson(contents)
	 if err != nil || atlas.Data.Status != 1{
	 	return
	 }
	  totalPage := editorialPersist.ParseAtlasPage(contents)
      for _,item := range atlas.Data.Data.List{
    	if item.Id == 0{
    		 continue
		}
    	ret.Requests = append(ret.Requests,engine.Request{
			Url: constant.BaseUrl+"/"+ strconv.FormatInt(item.Id,10)+"?groupid="+args.GroupId+"&",
			Method:"GET",
			Parser:engine.NewFuncParser(ParseEditorialPic,args.Title),
			Args:engine.RequestArgs{
				CategoryId:args.CategoryId,
				GroupId:args.GroupId,
				ItemId:item.Id,
				Data:item,
			},
		})
	}
	if  args.Page < totalPage{
		ret.Requests = append(ret.Requests,engine.Request{
			Url:constant.BaseUrl+"/group/"+args.GroupId,
			Method:"GET",
			Parser:engine.NewFuncParser(ParseEditorialAtlasTopic,args.Title),
			Args:engine.RequestArgs{
				CategoryId:args.CategoryId,
				Title:args.Title,
				NavId:args.NavId,
				GroupId:args.GroupId,
				Page:args.Page+1,
			},
		})
	}
	return ret
}
