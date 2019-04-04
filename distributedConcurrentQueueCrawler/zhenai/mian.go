package main

import (
	"crawler/distributedConcurrentQueueCrawler/config"
	"crawler/distributedConcurrentQueueCrawler/engine"
	"crawler/distributedConcurrentQueueCrawler/persist/client"
	"crawler/distributedConcurrentQueueCrawler/scheduler"
	"crawler/distributedConcurrentQueueCrawler/zhenai/parser"
	"fmt"
)

func main() {

	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.QueueConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 1000,
		ItemChan:  itemChan,
	}
	e.Run(engine.Request{
		Url:        config.CrawlerUrl,
		ParserFunc: parser.ParseCityList,
	})
}
