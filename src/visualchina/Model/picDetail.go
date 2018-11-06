package Model

import (
	"fmt"
	"libary/db"
	"logger"
)

type PicDetailDb struct {
	Id            int64
	PicId         int64
	Caption       string
	FileType      string
	Size          string
	StoreSize     string
	Specification string
	Cameraman     string
	Brand         string
	CopyRight     string
	Category      string

}

const PIC_DETAIL_TABLE = "visual_pic_detail"

func (pic *PicDetailDb) Save() (id int64, err error) {
	sql := fmt.Sprintf("INSERT INTO %s (pic_id,caption,file_type,size,store_size,specification,cameraman,brand,copyright,category)" +
		" VALUES(%d,\"%s\",'%s','%s','%s','%s','%s','%s','%s','%s')",
		PIC_DETAIL_TABLE,pic.PicId,pic.Caption,pic.FileType,pic.Size,pic.StoreSize,pic.Specification,pic.Cameraman,pic.Brand,pic.CopyRight,pic.Category)
	id, err = db.Insert(sql)
	if err != nil{
		logger.Error.Println("sql insert error:",err,"sql:",sql)
		return 0,err
	}
	return
}
