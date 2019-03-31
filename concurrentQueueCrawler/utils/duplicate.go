package utils

var visitedUrls = make(map[string]bool)

/**
去除重复的请求
*/
func IsDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
