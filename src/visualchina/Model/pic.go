package Model

import (
	"fmt"
	"libary/db"
	"logger"
)

type PicDb struct {
	Id            int64
	PicId         int64
	ImgId         int64
	ResId         string
	GroupId       int64
	RelateGroupId string
	TopicId       string
	CategoryId    int64
	Type          int
	Url           string
	Title         string
	ProviderId    int64
	NextPicId     int64
	ImageDate     int64
}

const PIC_TABLE = "visual_pic"

func (pic *PicDb) Save() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (img_id,pic_id,res_id,group_id,relate_group_id,topic_id,category_id,type,url,title,providerId," +
		"nextPicId,img_date) VALUES(%d,%d,'%s',%d,'%s','%s',%d,%d,'%s','%s',%d,%d,%d)",
		PIC_TABLE,pic.ImgId,pic.PicId,pic.ResId,pic.GroupId,pic.RelateGroupId,pic.TopicId,pic.CategoryId,pic.Type,pic.Url,pic.Title,pic.ProviderId,pic.NextPicId,pic.ImageDate)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}
