package persist

import (
	"hash/crc32"
	"encoding/hex"
	"visualchina/Model"
	"time"
	"strconv"
)

type NavStruct struct {
	Title string
	Url   string
	Type  string
}

type Editorial  struct{
	navChan chan NavStruct
	status chan bool
}

func (s *Editorial) NavRun(worker int)  {
	s.navChan = make(chan NavStruct,worker)
}

func (s *Editorial) NavWorker() chan NavStruct {
	return s.navChan
}

func (s *Editorial) NavSubmit(r NavStruct)  {
       s.navChan <- r
}

func (s *Editorial) NavSave() Model.NavDb {
	r := <-s.navChan
	url := hex.EncodeToString([]byte(r.Url))
	crcStr := crc32.ChecksumIEEE([]byte(url))
	t,_ := strconv.Atoi( r.Type)
	nav := Model.NavDb{
		Title:r.Title,
		Url:r.Url,
		Type: t,
		Crc32:crcStr,
		AddDate:time.Now().Unix(),
	}
	id, _ := nav.NavSave()
	nav.Id = id
	return  nav
}


