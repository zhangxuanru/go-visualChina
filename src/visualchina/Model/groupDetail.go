package Model

import (
	"fmt"
	"libary/db"
	"logger"
	"strings"
)

type GroupDetailDb struct {
	Id               int64
	GroupId          int64
	EqualwUrl        string
	EqualwImageId    int64
	EqualhUrl        string
	EqualhImageId    int64
	Width            int64
	Height           int64
	Url800           string
	Url800ImageId    int64
	GroupPics        string
	AddDate         int64
}

const GroupDetailTable  = "visual_group_detail"

func (group *GroupDetailDb) Save() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (group_id,equalw_url,equalw_image_id,equalh_url,equalh_image_id,width,height,url800,url800_image_id,groupPics,add_date) " +
		"VALUES(%d,'%s',%d,'%s',%d,%d,%d,'%s',%d,'%s',%d)",
		GroupDetailTable,group.GroupId,group.EqualwUrl,group.EqualwImageId,group.EqualhUrl,group.EqualhImageId,group.Width,group.Height,group.Url800,group.Url800ImageId,group.GroupPics,group.AddDate)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}

func  GetGroupDetailByGroupId(id int64,field []string) (map[string]string) {
	sql:= fmt.Sprintf("SELECT %s FROM %s WHERE group_id=%d",strings.Join(field,","), GroupDetailTable,id)
	r, _ := db.GetRow(sql)
	return r
}
