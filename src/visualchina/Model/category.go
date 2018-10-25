package Model

import (
	"fmt"
	"libary/db"
	"logger"
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

