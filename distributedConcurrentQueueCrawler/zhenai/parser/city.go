package parser

import (
	"crawler/distributedConcurrentQueueCrawler/engine"
	"regexp"
	"strings"
)

var profileRegex = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank"><span>([^<]+)</span></a>`)

var genderRegex = regexp.MustCompile(`<div class="content2">([^>]+)</div>`)

var cityUrlRegex = regexp.MustCompile(`href="(http://city.zhenai.com/[^"]+)"`)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRegex.FindAllSubmatch(contents, -1)
	subMatch := genderRegex.FindAllSubmatch(contents, -1)
	cityUrlMatches := cityUrlRegex.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for i, m := range matches {
		gender := strings.Split(string(subMatch[i][1]), ",")[0]
		// 去除空格
		gender = strings.Replace(gender, " ", "", -1)
		// 去除换行符
		gender = strings.Replace(gender, "\n", "", -1)

		//result.Items = append(
		//	result.Items, "User "+name+" "+gender)

		url := string(m[1])
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    url,
				Parser: NewProfileParse(string(m[2]), gender),
			})
	}

	for _, m := range cityUrlMatches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
			})
	}
	return result
}
