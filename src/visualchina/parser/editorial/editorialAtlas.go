package editorial

import (
	"engine"
	"fmt"
)

//图集页
func ParseEditorialAtlasTopic(contents []byte,url string,args engine.RequestArgs) engine.ParseResult {
	fmt.Println(string(contents))
	return engine.ParseResult{}
}

