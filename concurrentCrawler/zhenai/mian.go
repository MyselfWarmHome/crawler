package main

import (
	"crawler/concurrentCrawler/engine"
	"crawler/concurrentCrawler/scheduler"
	"crawler/concurrentCrawler/zhenai/parser"
)

func main() {

	//简单引擎示例
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://city.zhenai.com/",
	//	ParserFunc: parser.ParseCityList,
	//})

	//使用goroutine实现并发版的爬虫
	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkCount: 100}
	//e.Run(engine.Request{
	//	Url:        "http://city.zhenai.com/",
	//	ParserFunc: parser.ParseCityList,
	//})

	//使用队列+goroutine实现可控的并发版的爬虫
	e := engine.QueueConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100}
	e.Run(engine.Request{
		Url:        "http://city.zhenai.com/",
		ParserFunc: parser.ParseCityList,
	})
}
