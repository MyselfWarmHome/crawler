package config

const (
	//JsonRpc的协议
	NetProtocol = "tcp"

	//报错到ES使用的端口
	ItemSaverPort = 1234

	//es中的Index
	ElasticIndex = "dating_profile"

	//es中的Type
	ElasticType = "zhenai"

	//es的连接地址
	ElasticUrls = "http://47.101.50.131:9200"

	//爬虫爬取得初始地址
	CrawlerUrl = "http://city.zhenai.com/"

	//es中的报错暴露的service
	ItemSaverRpc = "ItemSaverService.SaveInfo"
)
