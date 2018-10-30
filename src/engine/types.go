package engine

type Request struct {
   Url string
   Method string
   Content string
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
	Id       int64
	Title    string
	Type     string
	Update   string
	Pid      int64
	GroupId  string
	NavId    int64
	Page     int64
	PerPage  int64
	CategoryId int64
	Data interface{}
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


