package engine

import (
	"crawler/concurrentQueueCrawler/utils"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	//将新的请求加入到队列中
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		//去除重复的请求
		if utils.IsDuplicate(r.Url) {
			continue
		}

		parserResult, err := worker(r)
		if err != nil {
			continue
		}
		//将爬取的新请求加入到队列中
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("获取的数据: %v", item)
		}
	}
}
