package Model

import (
	"fmt"
	"libary/db"
	"logger"
	"strings"
	"strconv"
)

type CategoryDb struct {
	Id              int64
	Type            string
	Code            string
	CategoryId      string
	CategoryPid     string
	CategoryName    string
	Pid             string
	CategoryUrl     string
	TotalCount      uint32
	Status          int
	AddDate         int64
}

const CATEGORY_TABLE = "visual_category"

func (category *CategoryDb) Save() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (type,code,category_id,category_pid,category_name,pid,category_url,total_count,add_date) VALUES('%s','%s','%s','%s','%s','%s','%s',%d,%d)",
		CATEGORY_TABLE,category.Type,category.Code,category.CategoryId,category.CategoryPid,category.CategoryName,category.Pid,category.CategoryUrl,category.TotalCount,category.AddDate)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}


func (category *CategoryDb) GetTagDataByCateId()(map[string]string) {
	sql:= fmt.Sprintf("SELECT category_id,category_name,category_pid,pid,type FROM %s WHERE category_id=%s",CATEGORY_TABLE,category.CategoryId)
	r, _ := db.GetRow(sql)
	return r
}


//根据分类ID查询所有二级分类
func GetSubCategoryList(categoryId int64) (mp map[int]CategoryDb) {
	mp = make(map[int]CategoryDb)
	sql:= fmt.Sprintf("SELECT id,category_id,category_name,pid FROM %s WHERE category_pid=%d",CATEGORY_TABLE,categoryId)
	result, _ := db.GetList(sql)
	if len(result) == 0{
	 	return mp
	}
	for k,val := range result{
		id,_ := strconv.ParseInt(val["id"],10,64)
		mp[k] = CategoryDb{
			CategoryId :  val["category_id"],
			CategoryName: val["category_name"],
			Pid: val["pid"],
			Id: id,
		}
	}
	return mp
}



//根据分类ID查询所有子分类【二级和三级分类】
func GetSubAllCategoryList(categoryId int64) (mp map[int]CategoryDb) {
	mp = make(map[int]CategoryDb)
	sql:= fmt.Sprintf("SELECT id,category_id,category_name,pid FROM %s WHERE category_pid=%d",CATEGORY_TABLE,categoryId)
	result, _ := db.GetList(sql)
	if len(result) == 0{
		return mp
	}
    catIdList := make([]string,0)
    j := 0
    for _,val := range result{
    	id,_ := strconv.ParseInt(val["id"],10,64)
		mp[j] = CategoryDb{
			 CategoryId :  val["category_id"],
			 CategoryName: val["category_name"],
			 Pid: val["pid"],
			 Id: id,
		}
		catIdList = append(catIdList,val["id"])
		j++
	}
	sql = fmt.Sprintf("SELECT id,category_id,category_name,pid FROM %s WHERE pid in (%s)",CATEGORY_TABLE,strings.Join(catIdList,","))
	result, _ = db.GetList(sql)
	if len(result) == 0{
		return mp
	}
	j++
	for _,val:= range result{
		id,_ := strconv.ParseInt(val["id"],10,64)
	    mp[j] = CategoryDb{
			CategoryId :  val["category_id"],
			CategoryName: val["category_name"],
			Pid: val["pid"],
			Id:id,
		}
		j++
	}
	return mp
}

//获取所有分类列表
func GetCategoryList(fields []string, where string)  {
     sql:= fmt.Sprintf("SELECT %s FROM %s WHERE %s",strings.Join(fields,","),CATEGORY_TABLE,where)
	 result, err := db.GetList(sql)

	fmt.Println(result)
	fmt.Println(err)

}
