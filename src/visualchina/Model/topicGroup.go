package Model

import (
	"fmt"
	"libary/db"
	"logger"
)

type TopicGroupDb struct {
	Id            int64
	TopicId       int64
	GroupId       int
	ImageId       int64
	Title         string
	Caption       string
	EqualwUrl     string
	EqualhUrl     string
	GroupPicsNum   int
	Width          int
	Height         int
	ImgDate       int64
	AddDate       int64
}

const TOPIC_GROUP_TABLE = "visual_topic_group"

func (topic *TopicGroupDb) Save() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (topic_id,group_id,image_id,title,caption,equalw_url,equalh_url,group_pics_num,width,height,img_date,add_date) " +
		"VALUES(%d,%d,%d,'%s','%s','%s','%s',%d,%d,%d,%d,%d)",
		TOPIC_GROUP_TABLE,topic.TopicId,topic.GroupId,topic.ImageId,topic.Title,topic.Caption,topic.EqualwUrl,topic.EqualhUrl,
		topic.GroupPicsNum,topic.Width,topic.Height,topic.ImgDate,topic.AddDate)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}
