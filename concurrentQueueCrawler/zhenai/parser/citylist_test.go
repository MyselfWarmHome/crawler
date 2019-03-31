package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {

	//content, _ := fetcher.Fetch("http://city.zhenai.com/")
	//fmt.Printf("%s\n",content)

	contents, err := ioutil.ReadFile("citylist_test.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("结果应该返回 %d个元素,但是返回了 %d个元素",
			resultSize, len(result.Requests))
	}

	expectedUrls := []string{
		"http://city.zhenai.com/aba", "http://city.zhenai.com/akesu", "http://city.zhenai.com/alashanmeng",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("希望获取的url #%d: %s;但是返回的结果是 %s",
				i, url, result.Requests[i].Url)
		}
	}

}
