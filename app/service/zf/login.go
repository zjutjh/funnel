package zf

import (
	"funnel/app/apis/zf"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils"
	"funnel/app/utils/security"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func login(username string, password string) (*model.User, error) {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	response, err := client.Get(zf.ZfLoginHome)
	if err != nil || response == nil || len(response.Cookies()) < 1 {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}
	JSESSIONID := response.Cookies()[0]

	request, _ := http.NewRequest("GET", zf.ZfLoginKaptcha, nil)
	request.AddCookie(JSESSIONID)
	response, _ = client.Do(request)
	s, _ := ioutil.ReadAll(response.Body)
	captcha, _ := utils.BreakCaptcha(s)

	request, _ = http.NewRequest("GET", zf.ZfLoginGetPublickey, nil)
	request.AddCookie(JSESSIONID)
	response, _ = client.Do(request)
	s, _ = ioutil.ReadAll(response.Body)

	encodePassword, _ := security.GetEncodePassword(s, []byte(password))
	data4 := url.Values{"yhm": {username}, "mm": {encodePassword}, "yzm": {captcha}}
	request, _ = http.NewRequest("POST", zf.ZfLoginHome, strings.NewReader(data4.Encode()))
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
	return service.SetUser(service.ZFPrefix, username, password, cookie)
}
