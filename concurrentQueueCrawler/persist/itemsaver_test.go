package persist

import (
	"context"
	"crawler/concurrentQueueCrawler/engine"
	"crawler/concurrentQueueCrawler/model"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSaveInfo(t *testing.T) {

	expected := engine.Item{
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

	//TODO start up ES from docker
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL("http://47.101.50.131:9200"))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = SaveInfo(client, expected, index)
	if err != nil {
		panic(err)
	}

	result, err := client.Get().Index(index).
		Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", result.Source)
	var actual engine.Item
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("Get %v ,but expected %v", actual, expected)
	}

}
