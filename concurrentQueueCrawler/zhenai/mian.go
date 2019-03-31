package main

import (
	"crawler/concurrentQueueCrawler/engine"
	"crawler/concurrentQueueCrawler/persist"
	"crawler/concurrentQueueCrawler/scheduler"
	"crawler/concurrentQueueCrawler/zhenai/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://city.zhenai.com/",
	//	ParserFunc: parser.ParseCityList,
	//})
	//
	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkCount: 100}
	//e.Run(engine.Request{
	//	Url:        "http://city.zhenai.com/",
	//	ParserFunc: parser.ParseCityList,
	//})

	itemChan, err := persist.ItemSaver("dating_profile", "http://47.101.50.131:9200")
	if err != nil {
		panic(err)
	}

	e := engine.QueueConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 1000,
		ItemChan:  itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://city.zhenai.com/",
		ParserFunc: parser.ParseCityList,
	})
}
