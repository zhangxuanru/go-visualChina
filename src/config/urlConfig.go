package config

import (
	"engine"
	"visualchina/parser"
)

type UrlStruct struct {
	 Url string
	 Name string
	 ParseFunc engine.ParserFunc
}

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
