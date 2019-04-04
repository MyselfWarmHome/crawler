package engine

import (
	"crawler/distributedConcurrentQueueCrawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	bodyByte, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(bodyByte, r.Url), nil
}
