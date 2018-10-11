package engine

type Request struct {
   Url string
   Parser Parser
}
type ParseResult struct {
	Requests []Request
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
}

type ParserFunc func(contents []byte, url string) ParseResult

type FuncParser struct {
	parser ParserFunc
	name string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return  f.Parse(contents,url)
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser:p,
		name:name,
	}
}


