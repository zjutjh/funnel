package libraryService

import (
	"funnel/app/apis/magicStrings"
	"net/url"
)

func genLoginData(username string, password string) url.Values {
	return url.Values{
		"__VIEWSTATE":          {magicStrings.Library__VIEWSTATE},
		"__VIEWSTATEGENERATOR": {magicStrings.Library__VIEWSTATEGENERATOR},
		"__EVENTVALIDATION":    {magicStrings.Library__EVENTVALIDATION},
		"TextBox1":             {username},
		"TextBox2":             {password},
		"ImageButton1.x":       {"29"},
		"ImageButton1.y":       {"8"}}
}
