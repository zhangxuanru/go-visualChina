package editorialPersist

import (
	"encoding/json"
	"strings"
	"regexp"
	"errors"
)

type TopicGroupJson struct {
	TopicInfo TopicInfo `json:"topicInfo"`
}

type TopicInfo struct {
	TopicImages TopicImages `json:"topicImages"`
}

type TopicImages struct {
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	Size     int         `json:"size"`
	OrderBy  interface{} `json:"orderBy"`
	StartRow int         `json:"startRow"`
	EndRow   int         `json:"endRow"`
	Total    int         `json:"total"`
	Pages    int         `json:"pages"`
	List     []TopicGroupList `json:"list"`
}

type TopicGroupList  struct {
	ID           int    `json:"id"`
	GroupID      string `json:"group_id"`
	Title        string `json:"title"`
	Caption      string `json:"caption"`
	EqualwURL    string `json:"equalw_url"`
	EqualhURL    string `json:"equalh_url"`
	ImgDate      string `json:"img_date"`
	GroupPicsNum int    `json:"group_pics_num"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

/*
 专题group的JSON字符串解析
*/
func ParseTopicGroupJson(contents []byte) (ret *TopicGroupJson,err error)  {
	var topic TopicGroupJson
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
	json.Unmarshal([]byte(content),&topic)
	return &topic,nil
}

