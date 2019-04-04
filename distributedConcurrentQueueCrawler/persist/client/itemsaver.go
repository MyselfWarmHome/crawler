package client

import (
	"crawler/distributedConcurrentQueueCrawler/config"
	"crawler/distributedConcurrentQueueCrawler/engine"
	"crawler/distributedConcurrentQueueCrawler/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	//获取JsonRpc的客户端
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: get item #%d: %v", itemCount, item)
			itemCount++

			//调用es的rpcServer
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v",
					item, err)
			}
		}
	}()
	return out, nil
}
