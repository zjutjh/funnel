package zfService

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"funnel/app/apis/oauth"
	"funnel/app/utils/fetch"
	"funnel/app/utils/security"
	"math/big"
	"net/url"
	"strconv"
)

func genTermExamInfoReqData(year string, term string, index int) url.Values {
	ksmc := []string{"",
		"EB4B492182673A1DE0550113465EF1CF",
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
		"queryModel.showCount": {"100"},
		"xsdm":                 {}}
}

func encryptPassword(password, modulus, exponent string) (string, error) {
	nString, _ := base64.StdEncoding.DecodeString(modulus)
	n, _ := new(big.Int).SetString(hex.EncodeToString(nString), 16)
	eString, _ := base64.StdEncoding.DecodeString(exponent)
	e, _ := strconv.ParseInt(hex.EncodeToString(eString), 16, 32)
	pub := rsa.PublicKey{E: int(e), N: n}
	cc, err := rsa.EncryptPKCS1v15(rand.Reader, &pub, []byte(password))
	return base64.StdEncoding.EncodeToString(cc), err
}

// func genLoginData(username, encryptedPwd string) url.Values {
// 	return url.Values{
// 		"yhm": {username},
// 		"mm":  {encryptedPwd}}
// }

// func genLoginData(username, password string, f fetch.Fetch) url.Values {
// 	s, _ := f.Get(zf.ZfLoginGetPublickey())
// 	encodePassword, _ := security.GetEncodePassword(s, []byte(password))
// 	return url.Values{
// 		"yhm": {username},
// 		"mm":  {encodePassword}}
// }

func genOauthLoginData(username, password, execution string, f *fetch.Fetch) url.Values {
	s, _ := f.Get(oauth.OauthLoginGetPublickey())

	encodePassword, _ := security.GetEncryptPassword(s, password)
	return url.Values{
		"username":   {username},
		"mobileCode": {},
		"password":   {encodePassword},
		"authcode":   {},
		"execution":  {execution},
		"_eventId":   {"submit"}}
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
