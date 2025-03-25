package zfService

import (
	"bytes"
	"encoding/json"
	"funnel/app/apis"
	"funnel/app/apis/oauth"
	"funnel/app/apis/zf"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/service/libraryService/request"
	"funnel/app/utils/fetch"
	"funnel/config"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strings"

	errorss "errors"
)

type captchaServerResponse struct {
	Status int    `json:"status"`
	Data   string `json:"msg"`
}

func login(username string, password string) (*model.User, error) {
	f := fetch.Fetch{}
	f.Init()
	_, err := f.Get(zf.ZfLoginHome())
	//if strings.Contains(err.Error(), "context deadline exceeded") {
	//	if strings.Compare(config.Redis.Get("zf_url").String(), "new") == 0 {
	//		config.Redis.Set("zf_url", "bk", 0)
	//	} else {
	//		config.Redis.Set("zf_url", "new", 0)
	//	}
	//}
	if err != nil {
		return nil, err
	}
	if len(f.Cookie) < 2 {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}
	var URL string
	if strings.Compare(config.Redis.Get("zf_url").String(), "new") == 0 {
		URL = apis.CAPTCHA_NEW_BREAKER_URL
	} else {
		URL = apis.CAPTCHA_BREAKER_URL
	}
	captcha, err := f.Get(URL + "?session=" + f.Cookie[0].Value + "&route=" + f.Cookie[1].Value)
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
	var sessionCookie *http.Cookie
	var routeCookie *http.Cookie
	for _, v := range f.Cookie {
		if v.Name == "JSESSIONID" {
			sessionCookie = v
		}
		if v.Name == "route" {
			routeCookie = v
		}
	}
	if sessionCookie == nil {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}
	return service.SetUser(service.ZFPrefix, username, password, sessionCookie, routeCookie)
}

func loginByOauth(username string, password string) (*model.User, error) {
	client := request.New()
	loginSuccess := errorss.New("login success")
	client.SetRedirectPolicy(
		resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
			// 自定义重定向处理逻辑
			if req.URL.String() == "http://www.gdjwjf.zjut.edu.cn/jwglxt/xtgl/login_slogin.html" {
				return loginSuccess
			}
			return nil
		}),
	)

	resp, err := client.R().Get(oauth.OauthLoginHome())
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return nil, err
	}
	hiddenInput := doc.Find("input[type=hidden][name=execution]")
	execution := hiddenInput.AttrOr("value", "")
	loginData := genOauthLoginData(username, password, execution, &client)
	resp, err = client.R().
		SetFormData(loginData).
		Post(resp.RawResponse.Request.URL.String())

	if !errorss.Is(err, loginSuccess) {
		if err != nil {
			return nil, err
		}
		err = CheckLogin(resp)
		return nil, err
	}

	cookies := resp.Cookies()

	var sessionCookie *http.Cookie
	var routeCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "JSESSIONID" {
			sessionCookie = cookie
		}
		if cookie.Name == "route" {
			routeCookie = cookie
		}
	}
	if sessionCookie == nil {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}

	// route 走 JF 时不需要route
	//if routeCookie == nil {
	//	return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	//}

	return service.SetUser(service.ZFPrefix, username, password, sessionCookie, routeCookie)
}
