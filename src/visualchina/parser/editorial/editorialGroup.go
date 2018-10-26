package editorial

import (
	"engine"
	"visualchina/Model"
	"strconv"
	"config/constant"
	"fmt"
)


//处理二级子导航页面
func ParseEditorialPageNavData(contents []byte,url string,args engine.RequestArgs) engine.ParseResult {
	//抓取图集start

	//抓取图集end
	return engine.ParseResult{}
}


//根据导航ID抓取group数据
func getGroupDataByNavId(id int64) (ret engine.ParseResult) {
	var categoryId int64
	fields :=[]string{"category_id","title"}
    result := Model.GetNavDataById(id,fields)
    if _, ok := result["category_id"]; !ok{
         return
	}
    categoryId, _ = strconv.ParseInt(result["category_id"], 10, 64)
	if categoryId == 0{
		return
	}
	 subCatList := Model.GetSubCategoryList(categoryId)
	 if len(subCatList) > 0{
	 	for _,val := range subCatList{
	 		subCategoryId,_:= strconv.ParseInt(val.CategoryId,10,64)
			ret.Requests = append(ret.Requests,engine.Request{
				Url:constant.GroupDataUrl,
				Method:"POST",
				Parser:engine.NewFuncParser(UpdateTotalNumberByCatId,val.CategoryName),
				Args:engine.RequestArgs{
					CategoryId: subCategoryId,
					NavId:id,
				},
				Content: fmt.Sprintf("key=%d&page=%d&per_page=%d&isEdit=1&timeliness=0",subCategoryId,1,5),
			})
		}
	 }
	ret.Requests = append(ret.Requests,engine.Request{
		Url:constant.GroupDataUrl,
		Method:"POST",
		Parser:engine.NewFuncParser(SaveGroup,result["Title"]),
		Args:engine.RequestArgs{
			CategoryId:categoryId,
			NavId:id,
		},
		Content: fmt.Sprintf("key=%d&page=%d&per_page=%d&isEdit=1&timeliness=0",categoryId,1,100),
	})

	return ret
}


//保存group数据
func SaveGroup(contents []byte,url string,args engine.RequestArgs) engine.ParseResult {
fmt.Println("SaveGroup:")
	fmt.Println(string(contents))
	fmt.Println(args)
	return engine.ParseResult{}
}


//获取分类的group总数， 更新visual_category表中的total_count字段，并根据args带的category_id 返回子栏目的request
func UpdateTotalNumberByCatId(contents []byte,url string,args engine.RequestArgs) engine.ParseResult {
    fmt.Println("UpdateTotalNumberByCatId:")
     fmt.Println(string(contents))
     fmt.Println(args)
	return engine.ParseResult{}
}
