package zfService

import (
	"funnel/app/apis/zf"
	"funnel/app/utils/fetch"
	"funnel/app/utils/security"
	"net/url"
)

func genTermRelatedInfoReqData(year string, term string) url.Values {
	return url.Values{"xnm": {year}, "xqm": {term}, "queryModel.showCount": {"1500"}}
}
func genLoginData(username, password string, f fetch.Fetch) url.Values {
	s, _ := f.Get(zf.ZfLoginGetPublickey())
	encodePassword, _ := security.GetEncodePassword(s, []byte(password))
	return url.Values{"yhm": {username}, "mm": {encodePassword}}
}
func genEmptyRoomReqData(year string, term string, campus string, week string, weekday string, classPeriod string) url.Values {
	return url.Values{
		"fwzt":                 {"cx"},
		"xnm":                  {year},
		"xqm":                  {term},
		"xqh_id":               {campus},
		"jyfs":                 {"0"},
		"zcd":                  {week},
		"xqj":                  {weekday},
		"jcd":                  {classPeriod},
		"queryModel.showCount": {"1500"}}
}
