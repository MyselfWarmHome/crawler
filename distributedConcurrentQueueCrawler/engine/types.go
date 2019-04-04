package engine

type ParserFunc func(contents []byte, url string) ParseResult

/**
请求实体
*/
type Request struct {
	Url        string
	ParserFunc ParserFunc
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

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
