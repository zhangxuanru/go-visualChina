package fetcher

import (
	"engine"
	"time"
	"visualchina/parser"
)

func Fetch(r engine.Request)(engine.ParseResult,error)  {
	time.Sleep(2*time.Second)

	result := engine.ParseResult{}
	request := engine.Request{
		Url:"http://www.baidu.com",
		Parser:engine.NewFuncParser(parser.ParseEditorial,"test1"),
	}
	request1 := engine.Request{
		Url:"http://www.sina.com",
		Parser:engine.NewFuncParser(parser.ParseEditorial,"test1"),
	}
	result.Requests = append(result.Requests,request,request1)
	return result,nil
}