package Model

import (
	"fmt"
	"libary/db"
	"logger"
)

type TopicCategoryDb struct {
	Id            int64
	CategoryId    int64
	CategoryName  string
	Link          string
	AddDate       int64
}

const TOPIC_CATEGORY_TABLE = "visual_topic_category"

func (topic *TopicCategoryDb) Save() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (category_id,category_name,link,add_date) VALUES(%d,'%s','%s',%d)",
		TOPIC_CATEGORY_TABLE,topic.CategoryId,topic.CategoryName,topic.Link,topic.AddDate)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}

func  GetTopicCategoryDataByCateId(categoryId int64) (row map[string]string) {
	sql:= fmt.Sprintf("SELECT category_id,category_name,link FROM %s WHERE category_id=%d",TOPIC_CATEGORY_TABLE,categoryId)
	r, _ := db.GetRow(sql)
	return r
}
