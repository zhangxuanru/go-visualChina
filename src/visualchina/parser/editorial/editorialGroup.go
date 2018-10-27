package editorial

import (
	"engine"
	"visualchina/Model"
	"strconv"
	"config/constant"
	"fmt"
	"visualchina/persist/editorialPersist"
	"net/http"
	"logger"
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
	 		subPid,_ := strconv.ParseInt(val.Pid,10,64)
			ret.Requests = append(ret.Requests,engine.Request{
				Url:constant.GroupDataUrl,
				Method:"POST",
				Parser:engine.NewFuncParser(UpdateTotalNumberByCatId,val.CategoryName),
				Args:engine.RequestArgs{
					CategoryId: subCategoryId,
					NavId:id,
					Pid:subPid,
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
			Title:result["Title"],
			NavId:id,
			Page:1,
			PerPage:100,
		},
		Content: fmt.Sprintf("key=%d&page=%d&per_page=%d&isEdit=1&timeliness=0",categoryId,1,100),
	})

	return ret
}


//保存group数据
func SaveGroup(contents []byte,url string,args engine.RequestArgs) (ret engine.ParseResult) {
	if len(contents) == 0{
          return
	}
	group:= editorialPersist.ParseGroupJson(contents)
	if group.Code != http.StatusOK || group.Status != 1{
		logger.Info.Println("SaveGroup json :",group)
		return
	}
	//这里根据keywords字段 按分类ID 拆成map, 然后各自加加，最后记录表中各分类共抓取多少数据
	for _,item := range group.Data.List{
         editorial.SaveGroup(item)
	}
	if (args.Page) *100 >= constant.MAXGroupData{
         return  engine.ParseResult{}
	}
	ret.Requests = append(ret.Requests,engine.Request{
		Url:constant.GroupDataUrl,
		Method:"POST",
		Parser:engine.NewFuncParser(SaveGroup,args.Title),
		Args:engine.RequestArgs{
			CategoryId:args.CategoryId,
			Title:args.Title,
			NavId:args.NavId,
			Page:args.Page+1,
			PerPage:100,
		},
		Content: fmt.Sprintf("key=%d&page=%d&per_page=%d&isEdit=1&timeliness=0",args.CategoryId,args.Page+1,100),
	})
	return  ret
}



//获取分类的group总数， 更新visual_category表中的total_count字段，并根据args带的category_id 返回子栏目的request
func UpdateTotalNumberByCatId(contents []byte,url string,args engine.RequestArgs) (ret engine.ParseResult) {
	group:= editorialPersist.ParseGroupJson(contents)
	if group.Code != http.StatusOK || group.Status != 1{
		logger.Info.Println("SaveGroup json :",group)
		return
	}
	editorial.UpdateCateGoryTotalNum(args.CategoryId,group)
	subCatList := Model.GetSubThreeCategoryList(args.Pid)
	if len(subCatList) > 0{
		for _,val := range subCatList{
			subCategoryId,_:= strconv.ParseInt(val.CategoryId,10,64)
			subPid,_ := strconv.ParseInt(val.Pid,10,64)
			ret.Requests = append(ret.Requests,engine.Request{
				Url:constant.GroupDataUrl,
				Method:"POST",
				Parser:engine.NewFuncParser(UpdateTotalNumberByCatId,val.CategoryName),
				Args:engine.RequestArgs{
					CategoryId: subCategoryId,
					NavId:args.NavId,
					Pid:subPid,
				},
				Content: fmt.Sprintf("key=%d&page=%d&per_page=%d&isEdit=1&timeliness=0",subCategoryId,1,5),
			})
		}
	}
	return ret
}
