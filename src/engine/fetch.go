package engine

import (
	"net/http"
	"logger"
	"io/ioutil"
	"os"
	"strings"
)

func FetchUrl(r Request) (ParseResult,error){
    if r.Method == "" || strings.ToUpper(r.Method) == "GET"{
          return FetchGet(r)
	  }
	 return FetchPost(r)
}



func FetchGet(r Request)(ParseResult,error)  {
	//test
	if r.Args.ItemId > 0{
		file, _ := os.Open("src/test/pic.html")
		all, _ := ioutil.ReadAll(file)
		return  r.Parser.Parse(all,r.Url,r.Args),nil
	}
	if r.Args.Id == 8{
		file, _ := os.Open("src/test/gundong.html")
		all, _ := ioutil.ReadAll(file)
		return  r.Parser.Parse(all,r.Url,r.Args),nil
	}
	if r.Args.UrlType == "topic"{
		file, _ := os.Open("src/test/topic.html")
		all, _ := ioutil.ReadAll(file)
		return  r.Parser.Parse(all,r.Url,r.Args),nil
	}
	if r.Args.UrlType == "topiccategory"{
		file, _ := os.Open("src/test/topic_category.html")
		all, _ := ioutil.ReadAll(file)
		return  r.Parser.Parse(all,r.Url,r.Args),nil
	}
	if r.Args.UrlType == "topicgroup"{
		file, _ := os.Open("src/test/topic_group.html")
		all, _ := ioutil.ReadAll(file)
		return  r.Parser.Parse(all,r.Url,r.Args),nil
	}
	if r.Args.Id > 0{
	    file, _ := os.Open("src/test/nav_view.html")
		all, _ := ioutil.ReadAll(file)
		return  r.Parser.Parse(all,r.Url,r.Args),nil
	}
	if len(r.Args.GroupId) > 0{
		file, _ := os.Open("src/test/group.html")
		all, _ := ioutil.ReadAll(file)
		return  r.Parser.Parse(all,r.Url,r.Args),nil
	}


	open, _ := os.Open("src/test/editorial.html")
	all, _ := ioutil.ReadAll(open)
	return  r.Parser.Parse(all,r.Url,r.Args),nil
   //test



	client := &http.Client{}
	request, e := http.NewRequest("GET", r.Url, nil)
	if e != nil{
		logger.Error.Println("HTTP GET 获取URL失败:",r.Url)
		return ParseResult{},nil
	}
	request.Header.Add("Accept-Language","zh-CN,zh;q=0.9")
	request.Header.Add("Cookie","acw_tc=b65cfd2215391390020578772e709f381d5adebcc60f8a8a1ef3d2c00079a9; _ga=GA1.2.106168907.1539139179; channel2=%2Fgroup%2F505378134%3Ffrom%3Drecommend; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%221665bd8140a2dc-049da456d42a04-333b5602-1296000-1665bd8140b4f9%22%2C%22%24device_id%22%3A%221665bd8140a2dc-049da456d42a04-333b5602-1296000-1665bd8140b4f9%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_referrer_host%22%3A%22%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%7D")
    request.Header.Add("Host","www.vcg.com")
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")

	response, err := client.Do(request)
	if err != nil{
		logger.Error.Println("HTTP GET 获取URL失败:",r.Url)
		return ParseResult{},nil
	}
	if response.StatusCode != http.StatusOK{
		logger.Error.Println("HTTP GET 获取URL状态码失败:",r.Url,response.StatusCode)
		return ParseResult{},nil
	}
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	return r.Parser.Parse(bytes,r.Url,r.Args),nil
}



func FetchPost(r Request)(ParseResult,error)  {
	//test
	   if r.Args.CategoryId > 0{
		   file, _ := os.Open("src/test/taglist.json")
		   defer file.Close()
		   bytes, _ := ioutil.ReadAll(file)
		   return  r.Parser.Parse(bytes,r.Url,r.Args),nil
	   }
	   return  ParseResult{},nil
	//test



	client := &http.Client{}
	request, e := http.NewRequest("POST", r.Url, strings.NewReader(r.Content))
	if e != nil{
		logger.Error.Println("HTTP POST 获取URL失败:",r.Url)
		return ParseResult{},nil
	}
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	request.Header.Add("Referer","https://www.vcg.com/editorial-channel-entertainment")
	request.Header.Add("Origin","https://www.vcg.com")
	request.Header.Add("Host","www.vcg.com")
	request.Header.Add("Content-Type","application/x-www-form-urlencoded")
	response, i := client.Do(request)
	if i != nil{
		logger.Error.Println("HTTP POST 获取URL失败:",r.Url)
		return ParseResult{},nil
	}
	if response.StatusCode != http.StatusOK{
		logger.Error.Println("HTTP GET 获取URL状态码失败:",r.Url,response.StatusCode)
		return ParseResult{},nil
	}
	defer response.Body.Close()
	body,_ := ioutil.ReadAll(response.Body)
	return r.Parser.Parse(body,r.Url,r.Args),nil
}

