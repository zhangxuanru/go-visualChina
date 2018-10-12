package persist

import (
	"fmt"
	"hash/crc32"
	"encoding/hex"
)

type NavStruct struct {
	Title string
	Url   string
}

type Editorial  struct{
	navChan chan NavStruct
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

func (s *Editorial) NavSave()  {
	r := <-s.navChan
	url := hex.EncodeToString([]byte(r.Url))
	crcStr := crc32.ChecksumIEEE([]byte(url))
    fmt.Println("navsave:",r,"url:",url,"crc32: ",crcStr)
}


