package main

import (
	"crawler/singleCrawler/engine"
	"crawler/singleCrawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://city.zhenai.com/",
		ParserFunc: parser.ParseCityList,
	})
}
