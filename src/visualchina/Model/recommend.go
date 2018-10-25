package Model

import (
	"fmt"
	"libary/db"
	"logger"
)

type RecommendDb struct {
	Id              int64
	CategoryId      string
	GroupId         string
	TopicId         string
	sort            int
	AddDate         int64
}

const RecommendTable = "visual_category_recommend"

func (recommend *RecommendDb) Save() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (category_id,group_id,topic_id,sort,add_date) VALUES('%s','%s','%s',%d,%d)",
		RecommendTable,recommend.CategoryId,recommend.GroupId,recommend.TopicId,recommend.sort, recommend.AddDate)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}

func (recommend *RecommendDb) GetRecommendDataById()(map[string]string) {
	sql := ""
	if len(recommend.GroupId) > 0{
		sql= fmt.Sprintf("SELECT category_id,group_id,topic_id,sort FROM %s WHERE category_id=%s AND group_id=%s",RecommendTable,
			recommend.CategoryId,recommend.GroupId)
	}
	if len(recommend.TopicId) > 0{
		sql= fmt.Sprintf("SELECT category_id,group_id,topic_id,sort FROM %s WHERE category_id=%s AND  topic_id=%s ",RecommendTable,
			recommend.CategoryId,recommend.TopicId)
	}
	r, _ := db.GetRow(sql)
	return r
}

