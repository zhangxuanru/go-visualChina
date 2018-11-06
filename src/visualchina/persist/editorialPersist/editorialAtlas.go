package editorialPersist

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"regexp"
	"errors"
	"strconv"
)

type AtlasJson struct {
	Data  DataList `json:"data"`
}

type DataList struct {
	Status int64  `json:"status"`
	Message string `json:"message"`
	Data   GroupData
}

type GroupData struct {
	ID int64    `json:"id"`
	GroupId string   `json:"group_id"`
	GroupTitle string `json:"group_title"`
	GroupCaption string `json:"group_caption"`
	GroupEditorName string `json:"group_editor_name"`
	TotalCount int64 `json:"total_count"`
	FirstPicId int64 `json:"first_pic_id"`
	List []PicList  `json:"list"`
}

type PicList struct {
	Id int64	   `json:"id"`
	ResId string	    `json:"res_id"`
	EqualwUrl string	  `json:"equalw_url"`
	EqualhUrl string	  `json:"equalh_url"`
	Url800 string   `json:"url800"`
	Title string  `json:"title"`
	PriceType string	`json:"price_type"`
	ImgDate string	`json:"img_date"`
	Width int	 `json:"width"`
	Height int  `json:"height"`
	Dlsize string	`json:"dlsize"`
	Caption string	`json:"caption"`
	ProviderId int64	`json:"providerId"`
	OneCategoryCn string  `json:"oneCategoryCn"`
	Category string `json:"category"`
	Tiffsize string `json:"tiffsize"`
}

/*
 group下的图集返回的JSON字符串解析
*/
func ParseAtlasJson(contents []byte) (ret *AtlasJson,err error)  {
	var atlas AtlasJson
	compile, err := regexp.Compile(`<script>window.__REDUX_STATE__(.*)</script>`)
	if err != nil{
	    return
	}
	find := compile.FindSubmatch(contents)
	if len(find)<2{
		return ret,errors.New("FindSubmatch not match data ")
	}
	content := strings.TrimSpace(string(find[1]))
	content = strings.TrimFunc(content, func(r rune) bool {
	      return r ==' ' || r ==';' || r == '='
	})
	json.Unmarshal([]byte(content),&atlas)
	return &atlas,nil
}

//获取总共有多少页
func ParseAtlasPage(contents []byte)(totalPage int64)  {
	reader := strings.NewReader(string(contents))
	document, _ := goquery.NewDocumentFromReader(reader)
	page := document.Find(".page").Eq(0).Text()
	page = strings.Trim(page," / ")
	totalPage,_ =  strconv.ParseInt(page,10,64)
	return  totalPage
}



