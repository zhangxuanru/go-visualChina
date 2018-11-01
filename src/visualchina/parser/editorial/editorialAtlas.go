package editorial

import (
	"engine"
	"visualchina/persist/editorialPersist"
	"fmt"
)

//图集页
func ParseEditorialAtlasTopic(contents []byte,url string,args engine.RequestArgs) (ret engine.ParseResult) {
	 atlas, err := editorialPersist.ParseAtlasJson(contents)
	 if err != nil || atlas.Data.Status != 1{
	 	return
	 }
	totalPage := editorialPersist.ParseAtlasPage(contents)
    groupId,page := args.GroupId,args.Page
    for _,item := range atlas.Data.Data.List{
		editorial.SaveAtlas(item)
	}

    //如果是最后一页，则不再继续请求图集的链接了
    if page == totalPage{
    	return
	}
	fmt.Printf("%+v----%d",atlas.Data.Message,totalPage)
	return engine.ParseResult{}
}




