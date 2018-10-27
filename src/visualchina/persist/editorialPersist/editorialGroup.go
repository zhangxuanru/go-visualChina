package editorialPersist

import "encoding/json"

type GroupJsonData struct {
	Status int64
	Code   int64
	Data  JsonData
}

type JsonData struct{
	TotalCount int64 `json:"total_count"`
	List []List
}

type List struct{
	Id int64
	GroupId string        `json:"group_id"`
	OneCategory int64     `json:"oneCategory"`
	Category interface{}       `json:"category"`
	Title  string
	Caption string
	Width int64
	Height int64
	EqualwUrl string           `json:"equalw_url"`
	EqualhUrl string           `json:"equalh_url"`
	FirstResId int64           `json:"firstResId"`
	ImgDate string             `json:"img_date"`
	GroupPicsNum int64         `json:"group_pics_num"`
	Keywords string            `json:"keywords"`
	Price string               `json:"price"`
	Url800 string              `json:"url800"`
	GroupPics []map[string]string      `json:"groupPics"`
	oneCategoryCn interface{}         `json:"oneCategoryCn"`
}

/*
 group返回的JSON字符串解析
*/
func ParseGroupJson(contents []byte) (GroupJsonData)  {
       var group GroupJsonData
	   json.Unmarshal(contents,&group)
       return  group
}

