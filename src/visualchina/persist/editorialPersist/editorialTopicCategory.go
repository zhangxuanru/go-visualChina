package editorialPersist

import (
	"encoding/json"
	"strings"
	"regexp"
	"errors"
)

type TopicCategoryJson struct {
	Data  CategoryDataInfo `json:"data"`
}

type CategoryDataInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data   CategoryData `json:"data"`
}

type CategoryData    struct {
	TotalCount int64 `json:"total_count"`
	List       []CategoryList  `json:"list"`
}

type CategoryList   struct {
	EqualhURL string `json:"equalh_url"`
	ImgDate   string `json:"img_date"`
	Link      string `json:"link"`
	Caption   string `json:"caption"`
	EqualwURL string `json:"equalw_url"`
	ID        int    `json:"id"`
	Cid       int
	Title     string `json:"title"`
}

/*
 专题分类的JSON字符串解析
*/
func ParseTopicCategoryJson(contents []byte) (ret *TopicCategoryJson,err error)  {
	var topic TopicCategoryJson
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

