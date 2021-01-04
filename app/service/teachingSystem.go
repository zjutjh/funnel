package service

import (
	"funnel/app/apis"
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/model"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type TeachingSystem struct {
}

func (t *TeachingSystem) GetClassTable(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(apis.ZfClassTable, year, term, stu)
}
func (t *TeachingSystem) GetExamInfo(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(apis.ZfExamInfo, year, term, stu)
}
func (t *TeachingSystem) GetScoreDetail(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(apis.ZfScoreDetail, year, term, stu)
}
func (t *TeachingSystem) GetScore(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(apis.ZfScore, year, term, stu)
}

func fetchTermRelatedInfo(requestUrl string, year string, term string, stu *model.User) (string, error) {

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

func (t *TeachingSystem) GetTrainingPrograms(stu *model.User) ([]byte, error) {

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	request, _ := http.NewRequest("GET", apis.ZfUserInfo, nil)
	request.AddCookie(&stu.Session)
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	s, exist := doc.Find("#pyfaxx_id").Attr("value")
	if exist {
		request, _ := http.NewRequest("GET", apis.ZfPY+s, nil)
		request.AddCookie(&stu.Session)
		res, _ := client.Do(request)
		s, _ := ioutil.ReadAll(res.Body)
		return s, nil
	}
	return nil, nil
}

func (t *TeachingSystem) GetEmptyRoomInfo(year string, term string, campus string, weekday string, week string, classPeriod string, stu *model.User) (string, error) {

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	requestData := url.Values{
		"fwzt":                 {"cx"},
		"xnm":                  {year},
		"xqm":                  {term},
		"xqh_id":               {campus},
		"jyfs":                 {"0"},
		"zcd":                  {week},
		"xqj":                  {weekday},
		"jcd":                  {classPeriod},
		"queryModel.showCount": {"1500"}}
	request, _ := http.NewRequest("POST", apis.ZfEmptyClassRoom, strings.NewReader(requestData.Encode()))
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

func (t *TeachingSystem) login(username string, password string) (*model.User, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	response, err := client.Get(apis.ZfLoginHome)
	if err != nil || response == nil || len(response.Cookies()) < 1 {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}
	JSESSIONID := response.Cookies()[0]

	request, _ := http.NewRequest("GET", apis.ZfLoginKaptcha, nil)
	request.AddCookie(JSESSIONID)
	response, _ = client.Do(request)
	s, _ := ioutil.ReadAll(response.Body)
	captcha, _ := helps.BreakCaptcha(s)

	request, _ = http.NewRequest("GET", apis.ZfLoginGetPublickey, nil)
	request.AddCookie(JSESSIONID)
	response, _ = client.Do(request)
	s, _ = ioutil.ReadAll(response.Body)

	encodePassword, _ := helps.GetEncodePassword(s, []byte(password))
	data4 := url.Values{"yhm": {username}, "mm": {encodePassword}, "yzm": {captcha}}
	request, _ = http.NewRequest("POST", apis.ZfLoginHome, strings.NewReader(data4.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(JSESSIONID)
	response, _ = client.Do(request)
	s, _ = ioutil.ReadAll(response.Body)

	if strings.Contains(string(s), "验证码输入错误") {
		return nil, errors.ERR_WRONG_Captcha
	}
	if strings.Contains(string(s), "用户名或密码不正确") {
		return nil, errors.ERR_WRONG_PASSWORD
	}
	var cookie *http.Cookie
	for _, v := range response.Cookies() {
		if v.Name == "JSESSIONID" {
			cookie = v
		}

	}
	if cookie == nil {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}
	return SetUser(ZFPrefix, username, password, cookie)
}

func (t *TeachingSystem) GetUser(username string, password string) (*model.User, error) {
	user, err := GetUser(ZFPrefix, username, password)
	if err != nil {
		return t.login(username, password)
	}
	return user, err
}
