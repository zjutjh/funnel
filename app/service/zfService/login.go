package zfService

import (
	"encoding/json"
	"funnel/app/apis"
	"funnel/app/apis/zf"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils/fetch"
	"net/http"
	"strings"
)

type captchaServerResponse struct {
	Status int    `json:"status"`
	Data   string `json:"msg"`
}

func login(username string, password string) (*model.User, error) {
	f := fetch.Fetch{}
	f.Init()
	_, err := f.Get(zf.ZfLoginHome())

	if err != nil {
		return nil, err
	}
	if len(f.Cookie) < 1 {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}

	captcha, err := f.Get(apis.CAPTCHA_NEW_BREAKER_URL + f.Cookie[0].Value)
	if err != nil {
		return nil, err
	}
	captchaRes := &captchaServerResponse{}
	_ = json.Unmarshal(captcha, captchaRes)
	if captchaRes.Status != 0 {
		return nil, errors.ERR_WRONG_Captcha
	}

	loginData := genLoginData(username, password, f)

	s, err := f.PostForm(zf.ZfLoginHome(), loginData)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(s), "请先滑动图片进行验证") {
		return nil, errors.ERR_WRONG_Captcha
	}
	if strings.Contains(string(s), "用户名或密码不正确") {
		return nil, errors.ERR_WRONG_PASSWORD
	}
	var cookie *http.Cookie
	for _, v := range f.Cookie {
		if v.Name == "JSESSIONID" {
			cookie = v
		}
	}
	if cookie == nil {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}
	return service.SetUser(service.ZFPrefix, username, password, cookie)
}
