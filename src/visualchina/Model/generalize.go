package Model

import (
	"fmt"
	"libary/db"
	"logger"
)

type GeneralizeDb struct {
	Id              int64
	CategoryId      string
	ImageId         string
	TopicId         string
	Gtype           string
	Title           string
	Src             string
	Link            string
	AddDate         int64
}

const GENERALIZE_TABLE = "visual_generalize"

func (g *GeneralizeDb) Save()(id int64, err error)  {
	sql := fmt.Sprintf("INSERT INTO %s (category_id,image_id,topic_id,gtype,title,src,link,add_date) VALUES(%s,%s,%s,%s,'%s','%s','%s',%d)",
		GENERALIZE_TABLE,g.CategoryId,g.ImageId,g.TopicId,g.Gtype,g.Title,g.Src,g.Link,g.AddDate)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}


func (g *GeneralizeDb) GetListByCateIdANDLink() (map[string]string) {
	sql:= fmt.Sprintf("SELECT category_id,image_id,topic_id,title,src,link,add_date FROM %s WHERE category_id=%s AND link='%s'",GENERALIZE_TABLE,g.CategoryId,g.Link)
	r, _ := db.GetRow(sql)
	return r
}

