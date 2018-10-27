package Model

import (
	"fmt"
	"libary/db"
	"logger"
	"strings"
)

type GroupDb struct {
	Id              int64
	GroupId         int64
	OneCategory     int64
	OneCategoryCn   string
	Category        string
	Title           string
	Caption         string
	FirstPicId      string
	GroupPicsNum    int64
	Keywords        string
	ImgDate         int64
	AddDate         int64
}

const GroupTable  = "visual_group"

func (group *GroupDb) Save() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (group_id,oneCategory, title,caption,first_pic_id,group_pics_num,keywords,img_date,add_date) " +
		"VALUES(%d,%d,'%s','%s','%s',%d,'%s',%d,%d)",
		GroupTable,group.GroupId,group.OneCategory,group.Title,group.Caption,group.FirstPicId,group.GroupPicsNum,group.Keywords,group.ImgDate,group.AddDate)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}


func  GetGroupDataByGroupId(id int64,field []string) (map[string]string) {
	sql:= fmt.Sprintf("SELECT %s FROM %s WHERE group_id=%d",strings.Join(field,","), GroupTable,id)
	r, _ := db.GetRow(sql)
	return r
}



