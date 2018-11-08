package editorialPersist

import (
	"encoding/json"
	"strings"
	"regexp"
	"errors"
)

type TopicJson struct {
	Data struct {
		Data  TopicData `json:"data"`
	} `json:"data"`
}

type TopicData struct {
	List [][]TopicList `json:"list"`
	Twelve []struct {
		ID   int64    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
		List []TopicList `json:"list"`
	} `json:"twelve"`
}

type TopicList struct  {
	ID            int         `json:"id"`
	Source        interface{} `json:"source"`
	InitTime      int64       `json:"initTime"`
	Type          int         `json:"type"`
	Mode          interface{} `json:"mode"`
	Cid           int         `json:"cid"`
	EqualWUrl     string      `json:"equalw_url"`
	Title         string      `json:"title"`
	Subhead       string      `json:"subhead"`
	Keywords      string      `json:"keywords"`
	Description   string      `json:"description"`
	TplID         int         `json:"tplId"`
	Logo          string      `json:"logo"`
	Cover         string      `json:"cover"`
	Banner        string      `json:"banner"`
	ShortName     string      `json:"shortName"`
	DisplayName   interface{} `json:"displayName"`
	CustomDomain  string      `json:"customDomain"`
	BeginTime     interface{} `json:"beginTime"`
	EndTime       interface{} `json:"endTime"`
	PublishTime   int64       `json:"publishTime"`
	UpdateTime    interface{} `json:"updateTime"`
	Status        int         `json:"status"`
	CreatedTime   string      `json:"createdTime"`
	CreatedBy     string      `json:"createdBy"`
	UpdatedTime   string      `json:"updatedTime"`
	UpdatedBy     string      `json:"updatedBy"`
	AssetType     int         `json:"assetType"`
	PreBanner     interface{} `json:"preBanner"`
	Timeliness    int         `json:"timeliness"`
	RunningStatus int         `json:"runningStatus"`
}

/*
 专题的JSON字符串解析
*/
func ParseTopicJson(contents []byte) (ret *TopicJson,err error)  {
	var topic TopicJson
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

