package persist

import (
	"context"
	"crawler/distributedConcurrentQueueCrawler/engine"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

/**
WARNING 分布式版本中废弃此方法
*/
func ItemSaver(index string, urls ...string) (chan engine.Item, error) {
	out := make(chan engine.Item)

	//连接es
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(urls...))
	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: get item #%d: %v", itemCount, item)
			itemCount++

			err := SaveInfo(client, item, index)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v",
					item, err)
			}
		}
	}()
	return out, nil
}

/**
使用es保存用户的信息
*/
func SaveInfo(client *elastic.Client,
	item engine.Item, index string) (err error) {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	//向ES中插入数据
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.
		Do(context.Background())
	if err != nil {
		return err
	}

	//fmt.Printf("%+v", response)
	return nil
}
