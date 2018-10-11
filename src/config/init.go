package config

import (
	"time"
	"fmt"
	"os"
	"log"
	"io"
	"flag"
	"engine"
	"visualchina/parser"
)

const (
	LOG_PATH_DIR = "src/log/"
)

var(
	Info *log.Logger
	Error *log.Logger
	FlagType int
	FlagAll  int
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
	flag.IntVar(&FlagType,"type", 0, "0:编辑图片,1:创意壁纸,2:创意图片,3:设计素材")
	flag.IntVar(&FlagAll,"all",0,"1:抓取所有,0:单个抓取")
	flag.Parse()
}

type UrlStruct struct {
	 Url string
	 Name string
	 ParseFunc engine.ParserFunc
}

/*
	    0:"https://www.vcg.com/editorial",         //编辑图片
		1:"https://www.vcg.com/sets/wallpaper",   //创意壁纸
		2:"https://www.vcg.com/creative",        //创意图片
		3:"https://www.vcg.com/design",         //设计素材
*/
func InitUrls() map[int]UrlStruct {
	urls := map[int] UrlStruct{
		0:{
			Url:"https://www.vcg.com/editorial",       //编辑图片
			Name:"editorial",
			ParseFunc:parser.ParseEditorial,
		},
		1:{
			Url:"https://www.vcg.com/sets/wallpaper",  //创意壁纸
			Name:"wallpaper",
			ParseFunc:parser.ParseEditorial,
		},
		2:{
			Url:"https://www.vcg.com/creative",        //创意图片
			Name:"creative",
			ParseFunc:parser.ParseEditorial,
		},
		3:{
			Url:"https://www.vcg.com/design",         //设计素材
			Name:"design",
			ParseFunc:parser.ParseEditorial,
		},
	}
	return urls
}
