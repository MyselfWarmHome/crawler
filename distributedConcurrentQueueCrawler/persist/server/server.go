package main

import (
	"crawler/distributedConcurrentQueueCrawler/config"
	"crawler/distributedConcurrentQueueCrawler/persist/rpc"
	"crawler/distributedConcurrentQueueCrawler/rpcsupport"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(ServeRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex, config.ElasticUrls))
}

func ServeRpc(host, index string, urls ...string) error {
	//连接es
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(urls...))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &rpc.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
