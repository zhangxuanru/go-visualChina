package Model

import (
	"fmt"
	"libary/db"
	"logger"
)

type CategoryGroupDb struct {
	Id              int64
	CategoryId      int64
	GroupId         int64
	Status          int
	AddDate         int64
}

const CATEGORY_GROUP_TABLE = "visual_category_group"

func (category *CategoryGroupDb) Save() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (category_id,group_id,status,add_date) VALUES(%d,%d,%d,%d)",
		CATEGORY_GROUP_TABLE,category.CategoryId,category.GroupId,category.Status,category.AddDate)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}

