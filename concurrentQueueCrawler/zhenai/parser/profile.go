package parser

import (
	"crawler/concurrentQueueCrawler/engine"
	"crawler/concurrentQueueCrawler/model"
	"regexp"
	"strconv"
)

var ageRegex = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
var marriageRegex = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)</div>`)
var xinZuoRegex = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+座)[^<]*</div>`)
var heightRegex = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d+])cm</div>`)
var weightRegex = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d+])kg</div>`)
var workAddressRegex = regexp.MustCompile(`<div class="m-btn purple"[^>]*>工作地:([^<]+)</div>`)
var inComeRegex = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)
var hoKouRegex = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯:([^<]+)</div>`)
var idUrlRegex = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string, gender string) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRegex))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRegex))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRegex))
	if err == nil {
		profile.Weight = weight
	}

	profile.Name = name
	profile.Gender = gender
	profile.Hokou = extractString(contents, hoKouRegex)
	profile.Income = extractString(contents, inComeRegex)
	profile.WorkAddress = extractString(contents, workAddressRegex)
	profile.Marriage = extractString(contents, marriageRegex)
	profile.XinZuo = extractString(contents, xinZuoRegex)

	return engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRegex),
				Payload: profile,
			},
		},
	}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
