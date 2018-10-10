package main

import (
	"time"
	"fmt"
	"os"
	"log"
	"io"
	"flag"
)

const (
	LOG_PATH_DIR = "src/log/"
)

var(
	Info *log.Logger
	Error *log.Logger
	flagType int
	flagAll  int
)

func init()  {
	year :=  time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	infoLog := fmt.Sprintf(LOG_PATH_DIR+"info-%d-%d-%d.log",year,month,day)
	errorLog := fmt.Sprintf(LOG_PATH_DIR+"error-%d-%d-%d.log",year,month,day)

	infoFile, _ := os.OpenFile(infoLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	errorFile, _ := os.OpenFile(errorLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)

	Info = log.New(infoFile,"Info: ",log.LstdFlags | log.Lshortfile | log.LUTC)
	Error = log.New(io.MultiWriter(os.Stderr,errorFile),"error: ",log.LstdFlags | log.Lshortfile | log.LUTC)

	//获取命令行参数
	flag.IntVar(&flagType,"type", 0, "0:编辑图片,1:创意壁纸,2:创意图片,3:设计素材")
	flag.IntVar(&flagAll,"all",0,"1:抓取所有,0:单个抓取")
	flag.Parse()
}


func initUlrs() map[int]string {
	urls := map[int]string{
		0:"https://www.vcg.com/editorial",         //编辑图片
		1:"https://www.vcg.com/sets/wallpaper",   //创意壁纸
		2:"https://www.vcg.com/creative",        //创意图片
		3:"https://www.vcg.com/design",         //设计素材
	}
	return urls
}

