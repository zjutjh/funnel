package model

import (
	"funnel/app/apis"
	"funnel/app/errors"
	"funnel/app/helps"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type TeachingSystem struct {
	RootUrl       string
	ClassTableUrl string
}

func (t *TeachingSystem) GetClassTable(stu *ZFUser, year string, term string) (string,error) {
	return fetchTermRelatedInfo(apis.ZfClassTable, year, term, stu)
}
func (t *TeachingSystem) GetExamInfo(stu *ZFUser, year string, term string) (string,error) {
	return fetchTermRelatedInfo(apis.ZfExamInfo, year, term, stu)
}
func (t *TeachingSystem) GetScoreDetail(stu *ZFUser, year string, term string) (string,error) {
	return fetchTermRelatedInfo(apis.ZfScoreDetail, year, term, stu)
}
func (t *TeachingSystem) GetScore(stu *ZFUser, year string, term string) (string,error) {
	return fetchTermRelatedInfo(apis.ZfScore, year, term, stu)
}

func fetchTermRelatedInfo(requestUrl string, year string, term string, stu *ZFUser) (string, error) {

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	requestData := url.Values{"xnm": {year}, "xqm": {term}, "queryModel.showCount": {"1500"}}
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(requestData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(&stu.Session)
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	s, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(s), nil
}

func (t *TeachingSystem) Login(stu *ZFUser) error {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	url0 := apis.ZfLoginHome
	response, _ := client.Get(url0)
	JSESSIONID := response.Cookies()[0]

	publicKeyUrl := apis.ZfLoginGetPublickey
	request, _ := http.NewRequest("GET", publicKeyUrl, nil)
	request.AddCookie(JSESSIONID)
	response, _ = client.Do(request)
	s, _ := ioutil.ReadAll(response.Body)
	encodePassword, _ := helps.GetEncodePassword(s, []byte(stu.Password))

	url2 := apis.ZfLoginKaptcha
	request, _ = http.NewRequest("GET", url2, nil)
	request.AddCookie(JSESSIONID)
	response, _ = client.Do(request)
	s, _ = ioutil.ReadAll(response.Body)
	captcha, _ := helps.BreakCaptcha(s)

	url4 := apis.ZfLoginHome
	data4 := url.Values{"yhm": {stu.Username}, "mm": {encodePassword}, "yzm": {captcha}}
	request, _ = http.NewRequest("POST", url4, strings.NewReader(data4.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(JSESSIONID)

	response, _ = client.Do(request)
	s, _ = ioutil.ReadAll(response.Body)

	if strings.Contains(string(s), "验证码") {
		return errors.ERR_UNKNOWN_LOGIN_ERROR
	}
	if strings.Contains(string(s), "用户名或密码不正确") {
		return errors.ERR_WRONG_PASSWORD
	}
	if len(response.Cookies()) < 2 {
		return errors.ERR_UNKNOWN_LOGIN_ERROR
	}

	JSESSIONID = response.Cookies()[1]
	stu.Session = *JSESSIONID
	return nil
}
