package parser

import (
	"crawler/concurrentCrawler/engine"
	"regexp"
	"strings"
)

var profileRegex = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank"><span>([^<]+)</span></a>`)

var genderRegex = regexp.MustCompile(`<div class="content2">([^>]+)</div>`)

var cityUrlRegex = regexp.MustCompile(`href="(http://city.zhenai.com/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRegex.FindAllSubmatch(contents, -1)
	subMatch := genderRegex.FindAllSubmatch(contents, -1)
	cityUrlMatches := cityUrlRegex.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for i, m := range matches {
		name := string(m[2])
		gender := strings.Split(string(subMatch[i][1]), ",")[0]
		// 去除空格
		gender = strings.Replace(gender, " ", "", -1)
		// 去除换行符
		gender = strings.Replace(gender, "\n", "", -1)
		result.Items = append(
			result.Items, "User "+name+" "+gender)

		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: func(contents []byte) engine.ParseResult {
					return ParseProfile(contents, name, gender)
				},
			})
	}

	for _, m := range cityUrlMatches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}
	return result
}
