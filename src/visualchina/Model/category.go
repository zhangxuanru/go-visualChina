package Model

import (
	"fmt"
	"libary/db"
	"logger"
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



//根据PID查询所有子分类【三级分类】
func GetSubThreeCategoryList(Pid int64) (mp map[int]CategoryDb) {
	mp = make(map[int]CategoryDb)
	sql:= fmt.Sprintf("SELECT id,category_id,category_name,pid FROM %s WHERE pid=%d",CATEGORY_TABLE,Pid)
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



//更新分类总数
func UpdateCateGoryTotalNum(categoryId int64,total int64)(count int64,err error){
	sql := fmt.Sprintf("UPDATE %s SET total_count=%d WHERE category_id=%d",CATEGORY_TABLE,total,categoryId)
	count, err = db.UpdateSql(sql)
	return
}

