package engine

type Request struct {
   Url string
   Args RequestArgs
   Parser Parser
}

type Item struct {
	Title   string
	Url     string
	Type    string
	Data   interface{}
}

type RequestArgs struct {
	Title string
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}


type Parser interface {
	Parse(contents []byte, url string,args RequestArgs) ParseResult
	GetName() string
}

type ParserFunc func(contents []byte, url string,args RequestArgs) ParseResult

type FuncParser struct {
	parser ParserFunc
	name string
}

func (f *FuncParser) Parse(contents []byte, url string,args RequestArgs) ParseResult {
	return  f.parser(contents,url,args)
}

func (f *FuncParser) GetName() string {
	return f.name
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser:p,
		name:name,
	}
}


