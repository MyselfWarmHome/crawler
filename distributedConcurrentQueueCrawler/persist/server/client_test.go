package main

import (
	"crawler/distributedConcurrentQueueCrawler/engine"
	"crawler/distributedConcurrentQueueCrawler/model"
	"crawler/distributedConcurrentQueueCrawler/rpcsupport"
	"testing"
	"time"
)

func TestItemServer(t *testing.T) {
	const host = ":1234"
	go ServeRpc(host, "test1", "http://47.101.50.131:9200")

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1303084905",
		Type: "zhenai",
		Id:   "e32fafrarefregrtgtr",
		Payload: model.Profile{
			Age:         34,
			Height:      162,
			Weight:      57,
			Income:      "3000-5000元",
			Gender:      "女",
			Name:        "安静的雨",
			XinZuo:      "天蝎座",
			Marriage:    "未婚",
			Hokou:       "重庆",
			WorkAddress: "北京",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.SaveInfo", item, &result)
	if err != nil && result != "ok" {
		t.Errorf("result: %s; error: %s", result, err)
	}

}
