package Model

import (
	"fmt"
	"libary/db"
	"logger"
)

type TopicDb struct {
	TopicId          int
	ImageId          int64
	Type             int
	CategoryId       int
	Title            string
	Keywords         string
	Description      string
	EqualwUrl        string
	Cover            string
	IsBanner         int
	Status           int
	CreatedYear      int
	CreatedTime      int64
	UpdatedTime      int64
	AddDate          int64
}

const TOPIC_TABLE = "visual_topic"

func (topic *TopicDb) Save() (id int64, err error) {
	sql := fmt.Sprintf("REPLACE INTO %s (topic_id,image_id,type,category_id,title,keywords,description,equalw_url," +
		"cover,is_banner,status,created_year,created_time,updated_time,add_date) VALUES(%d,%d,%d,%d,'%s','%s','%s','%s','%s',%d,%d,%d,%d,%d,%d)",
		TOPIC_TABLE,topic.TopicId,topic.ImageId,topic.Type,topic.CategoryId,topic.Title,topic.Keywords,topic.Description,topic.EqualwUrl,
		topic.Cover,topic.IsBanner,topic.Status,topic.CreatedYear,topic.CreatedTime,topic.UpdatedTime,topic.AddDate)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}

func GetTopicDataById(topicId int64) (row map[string]string) {
	sql:= fmt.Sprintf("SELECT topic_id,title FROM %s WHERE topic_id=%d",TOPIC_TABLE,topicId)
	r, _ := db.GetRow(sql)
	return r
}
