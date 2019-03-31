package engine

import (
	"crawler/singleCrawler/fetcher"
	"crawler/singleCrawler/utils"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	//将新的请求加入到队列中
	for _, seed := range seeds {
		if utils.IsDuplicate(seed.Url) {
			continue
		}
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		//去除重复的请求
		if utils.IsDuplicate(r.Url) {
			continue
		}

		log.Printf("Fetching %s", r.Url)
		bodyByte, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v",
				r.Url, err)
			continue
		}

		//将爬取的新请求加入到队列中
		parserResult := r.ParserFunc(bodyByte)
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("获取的数据: %v", item)
		}
	}
}
