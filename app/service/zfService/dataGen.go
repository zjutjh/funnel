package zfService

import (
	"funnel/app/apis/zf"
	"funnel/app/utils/fetch"
	"funnel/app/utils/security"
	"net/url"
)

func genTermExamInfoReqData(year string, term string, index int) url.Values {
	ksmc := []string{"EB4B492182673A1DE0550113465EF1CF",
		"EB4ADB3912953991E0550113465EF1CF",
		"EB4ADB39129D3991E0550113465EF1CF",
		"EB4ADB3912993991E0550113465EF1CF",
		"E9B7D38A2E1907CFE0550113465EF1CF",
		"EA7B8C23F41E2384E0550113465EF1CF"}
	return url.Values{
		"xnm":                  {year},
		"xqm":                  {term},
		"ksmcdmb_id":           {ksmc[index]},
		"kzlx":                 {"ck"},
		"queryModel.showCount": {"100"}}
}
func genTermRelatedInfoReqData(year string, term string) url.Values {
	return url.Values{
		"xnm":                  {year},
		"xqm":                  {term},
		"kzlx":                 {"ck"},
		"queryModel.showCount": {"100"}}
}
func genLoginData(username, password string, f fetch.Fetch) url.Values {
	s, _ := f.Get(zf.ZfLoginGetPublickey())
	encodePassword, _ := security.GetEncodePassword(s, []byte(password))
	return url.Values{
		"yhm": {username},
		"mm":  {encodePassword}}
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
