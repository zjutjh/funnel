package comm

import "regexp"

func ExtractJsonStr(jsonStr, jsonKey string) string {
	re := regexp.MustCompile(jsonKey + `:'([^']+)'`)
	matches := re.FindStringSubmatch(jsonStr)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
