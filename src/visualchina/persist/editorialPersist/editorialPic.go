package editorialPersist

import (
	"encoding/json"
	"strings"
	"regexp"
	"errors"
)

type PicJson struct {
	Data  Info `json:"data"`
}

type Info struct {
	PicInfo     picInfo             `json:"picInfo"`
	GroupInfo   groupInfo           `json:"groupInfo"`
	RelateGroup []relateGroupInfo   `json:"relateGroup"`
	Topic       []topicInfo         `json:"topic"`
}

type picInfo struct {
	Id                int64
	ImgId             int64
	ResId             string  `json:"res_id"`
	Type              int
	Title             string
	Caption           string
	Url               string
	FileType          string  `json:"file_type"`
	Size              string
	StoreSize         string  `json:"store_size"`
	Cameraman         string  `json:"cameraman"`
	CopyRight         string  `json:"copyright"`
	Specification     string
	License           string
	ProviderId        int64
	Brand             string
	Category          string
	NextPicId         int64
	ImageDate         string
}

type groupInfo struct {
	GroupId string
	GroupTitle string
	GroupExplain string
}

type relateGroupInfo struct {
	GroupId string            `json:"groupId"`
	GroupTitle string         `json:"groupTitle"`
	GroupExplain string       `json:"groupExplain"`
}

type topicInfo struct {
	TopicId int64
	TopicName string
}


/*
 图片详情页的JSON字符串解析
*/
func ParsePicJson(contents []byte) (ret *PicJson,err error)  {
	var pic PicJson
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
	json.Unmarshal([]byte(content),&pic)
	return &pic,nil
}

