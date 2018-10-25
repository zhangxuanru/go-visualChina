package Model

import (
	"fmt"
	"libary/db"
	"logger"
)

type NavDb struct {
	Id              int64
	Title           string
	CategoryId      int64
	Type            int
    Url             string
	GrabUrl         string
    Crc32           uint32
    AddDate         int64
	ExecDate        int64
}

const table  = "visual_nav"

func (nav *NavDb) NavSave() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (title,category_id,type,url,grab_url,grab_url_crc32,exec_date,add_date) VALUES('%s',%d,%d,'%s','%s',%d,%d,%d)",
		table,nav.Title,nav.CategoryId,nav.Type,nav.Url,nav.GrabUrl,nav.Crc32,nav.ExecDate,nav.AddDate)
	 id, err = db.Insert(sql)
	 if err != nil{
	 	 logger.Error.Println("sql insert error:",err,"sql:",sql)
	 	 return 0,err
	 }
	 return
}

func (nav *NavDb) GetNavDataByCrc32(crc32 uint32) (map[string]string) {
	sql:= fmt.Sprintf("SELECT id,type,category_id,url,grab_url FROM %s WHERE grab_url_crc32=%d",table,crc32)
	r, _ := db.GetRow(sql)
	return r
}

func (nav *NavDb) UpdateNavExecTimeById(id int64,execDate int64) (count int64,err error) {
       sql := fmt.Sprintf("UPDATE %s SET exec_date=%d WHERE id=%d",table,execDate,id)
	   count, err = db.UpdateSql(sql)
	   return
}

//根据标题更新分类ID
func (nav *NavDb) UpdateNavCategoryIdByTitle() (count int64,err error) {
	sql := fmt.Sprintf("UPDATE %s SET category_id=%d WHERE title='%s'",table,nav.CategoryId,nav.Title)
	count, err = db.UpdateSql(sql)
	return
}



