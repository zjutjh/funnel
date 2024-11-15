package libraryService

import (
	"net/url"
	"regexp"
	"strings"
)

// Result 交由Resty实现response的自动反序列化
type Result struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrCode   int         `json:"errCode"`
	ErrorCode interface{} `json:"errorCode"`
	Data      struct {
		SearchResult []interface{} `json:"searchResult"`
		NumFound     int           `json:"numFound"`
	} `json:"data"`
}

// GetRedirectLocation 获取js中的重定向连接
func GetRedirectLocation(html string) string {
	re := regexp.MustCompile(`window\.location\s*=\s*"([^"]+)"`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		redirectURL := matches[1]
		// 去除转义
		redirectURL = regexp.MustCompile(`\\/`).ReplaceAllString(redirectURL, "/")

		// 处理unicode
		redirectURL = strings.Replace(redirectURL, `\u0026`, "&", -1)

		// 从redirectURL解析出refer
		u, err := url.Parse(redirectURL)
		if err != nil {
			return ""
		}
		queries, _ := url.ParseQuery(u.RawQuery)
		refer := queries["refer"][0]

		return refer
	}
	return ""
}
