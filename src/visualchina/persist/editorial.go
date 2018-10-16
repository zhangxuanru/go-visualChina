package persist

import (
	"hash/crc32"
	"encoding/hex"
	"visualchina/Model"
	"time"
	"strconv"
)


type Editorial  struct{
	NavDb  chan Model.NavDb
	status chan bool
}


func (s *Editorial) NavSave(r Model.NavDb) Model.NavDb {
	url := hex.EncodeToString([]byte(r.Url))
	crcStr := crc32.ChecksumIEEE([]byte(url))
	nav := Model.NavDb{
		Title: r.Title,
		Pid: r.Pid,
		Url: r.Url,
		Type: r.Type,
		Crc32: crcStr,
		AddDate:time.Now().Unix(),
	}
	 row := nav.GetNavDataByCrc32(crcStr,r.Pid)
	if len(row) >0 {
		id, _ := strconv.ParseInt(row["id"], 10, 64)
		nav.Id = id
		return nav
	}
	id, _ := nav.NavSave()
	nav.Id = id
	return  nav
}


