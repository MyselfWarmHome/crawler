package rpc

import (
	"crawler/distributedConcurrentQueueCrawler/engine"
	"crawler/distributedConcurrentQueueCrawler/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (e *ItemSaverService) SaveInfo(item engine.Item, result *string) error {
	err := persist.SaveInfo(e.Client, item, e.Index)
	log.Printf("Item %v saved", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Item: %v saving error: %v", item, err)
	}
	return err
}
