package engine

type ParserFunc func(contents []byte, url string) ParseResult

/**
序列化和反序列化 + 下一次Parser的接口
*/
type Parser interface {
	Parser(contents []byte, url string) ParseResult
	Serialize() (name string, gender string, args interface{})
}

/**
请求实体
*/
type Request struct {
	Url    string
	Parser Parser
}

/**
请求结果实体
*/
type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

/**
空的Parser的序列化与反序列化
*/
type NilParse struct {
}

func (NilParse) Parser(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParse) Serialize() (name string, gender string, args interface{}) {
	return "nilParser", "nilParser", nil
}

/**
包装具体的爬虫Parser
*/
type FuncParser struct {
	parser ParserFunc
	name   string //函数的名称
}

func (f *FuncParser) Parser(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, _ string, args interface{}) {
	return f.name, "", nil
}

/**
工厂模式生成FuncParser
*/
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
