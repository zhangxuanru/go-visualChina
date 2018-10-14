package Model

import (
	"fmt"
	"libary/db"
	"logger"
)

type NavDb struct {
	Id              int64
    Title           string
    Url             string
    Crc32           uint32
    AddDate         int64
	LastCrawlTime   int64
	Type            int
}

const table  = "visual_page_nav"

func (nav *NavDb) NavSave() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (type,nav_title,nav_url,nav_crc32,add_date,last_crawl_time) VALUES(%d,'%s','%s',%d,%d,%d)",
		table,nav.Type,nav.Title,nav.Url,nav.Crc32,nav.AddDate,nav.LastCrawlTime)
	 id, err = db.Insert(sql)
	 if err != nil{
	 	 logger.Error.Println("sql insert error:",err,"sql:",sql)
	 	 return 0,err
	 }
	 return
}
