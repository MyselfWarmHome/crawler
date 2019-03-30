package parser

import (
	"crawler/concurrentCrawler/engine"
	"fmt"
	"regexp"
)

const cityListRegex = `<a href="(http://city.zhenai.com/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRegex)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	//limit := 10
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		//limit--
		//if limit == 0 {
		//	break
		//}
	}

	fmt.Printf("匹配的个数: %d 个\n", len(matches))
	return result
}
