package zfService

import (
	"fmt"
	"funnel/app/apis/oauth"
	"funnel/app/apis/zf"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/service/zfService/loginTokenManager"
	"funnel/app/utils/fetch"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

type captchaServerResponse struct {
	Status int    `json:"status"`
	Data   string `json:"msg"`
}

func login(username string, password string) (*model.User, error) {
	ltk, err := loginTokenManager.LoginTkMgr.GetAToken()
	if err != nil {
		return nil, err
	}
	session, err := submitLogin(username, password, *ltk)
	if err != nil {
		return nil, err
	}
	return service.SetUser(service.ZFPrefix, username, password,
		&http.Cookie{Name: "JSESSIONID", Value: session},
		&http.Cookie{Name: "route", Value: ltk.Route})
}

func submitLogin(username, password string, loginToken loginTokenManager.LoginToken) (string, error) {
	loginURL := zf.ZfLoginHome()

	// 密码加密
	encPwd, err := encryptPassword(password, loginToken.CryptoModulus, loginToken.CryptoExponent)
	if err != nil {
		return "", fmt.Errorf("cannot encrypt password: %w", err)
	}

	// 构造 resty 客户端
	client := resty.New().
		SetTimeout(5 * time.Second).
		SetRedirectPolicy(resty.NoRedirectPolicy()) // 禁止自动重定向

	// 构造 Cookie 字符串
	cookies := fmt.Sprintf("JSESSIONID=%s; route=%s",
		loginToken.JSESSIONID,
		loginToken.Route,
	)

	// 发送 POST 请求
	resp, err := client.R().
		SetHeader("Cookie", cookies).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"yhm": username,
			"mm":  encPwd,
		}).Post(loginURL)
	// 处理响应 由于 302 会被认为是错误所以需要特殊处理
	if err != nil && resp.StatusCode() != http.StatusFound {
		return "", err
	}
	if resp.StatusCode() != http.StatusFound {
		// 正方的登录成功响应是 302 重定向 失败反而是 200
		// 在第一次重定向就能拿到新的 JSESSIONID 字段 后面的请求不必进行
		if resp.StatusCode() == http.StatusOK {
			// 如果是 200 则需要判断是什么原因导致的失败
			if strings.Contains(resp.String(), "用户名或密码不正确") {
				return "", errors.ERR_WRONG_PASSWORD
			}
			if strings.Contains(resp.String(), "请先滑动图片进行验证") {
				return "", errors.ERR_WRONG_Captcha
			}
			return "", errors.ERR_UNKNOWN_LOGIN_ERROR
		}
		return "", errors.ERR_UNKNOWN_LOGIN_ERROR
	}

	// 提取成功登录后更换的 JSESSIONID
	var sessionCookie string
	for _, c := range resp.Cookies() {
		if c.Name == "JSESSIONID" {
			sessionCookie = c.Value
		}
	}
	if sessionCookie == "" {
		return "", fmt.Errorf("login failed: JSESSIONID not found in response")
	}

	return sessionCookie, nil
}

// func login(username string, password string) (*model.User, error) {
// 	f := fetch.Fetch{}
// 	f.Init()
// 	_, err := f.Get(zf.ZfLoginHome())
// 	//if strings.Contains(err.Error(), "context deadline exceeded") {
// 	//	if strings.Compare(config.Redis.Get("zf_url").String(), "new") == 0 {
// 	//		config.Redis.Set("zf_url", "bk", 0)
// 	//	} else {
// 	//		config.Redis.Set("zf_url", "new", 0)
// 	//	}
// 	//}
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(f.Cookie) < 2 {
// 		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
// 	}
// 	var URL string
// 	if strings.Compare(config.Redis.Get("zf_url").String(), "new") == 0 {
// 		URL = apis.CAPTCHA_NEW_BREAKER_URL
// 	} else {
// 		URL = apis.CAPTCHA_BREAKER_URL
// 	}
// 	captcha, err := f.Get(URL + "?session=" + f.Cookie[0].Value + "&route=" + f.Cookie[1].Value)
// 	if err != nil {
// 		return nil, err
// 	}
// 	captchaRes := &captchaServerResponse{}
// 	_ = json.Unmarshal(captcha, captchaRes)
// 	if captchaRes.Status != 0 {
// 		return nil, errors.ERR_WRONG_Captcha
// 	}
// 	loginData := genLoginData(username, password, f)
// 	s, err := f.PostForm(zf.ZfLoginHome(), loginData)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if strings.Contains(string(s), "请先滑动图片进行验证") {
// 		return nil, errors.ERR_WRONG_Captcha
// 	}
// 	if strings.Contains(string(s), "用户名或密码不正确") {
// 		return nil, errors.ERR_WRONG_PASSWORD
// 	}
// 	var sessionCookie *http.Cookie
// 	var routeCookie *http.Cookie
// 	for _, v := range f.Cookie {
// 		if v.Name == "JSESSIONID" {
// 			sessionCookie = v
// 		}
// 		if v.Name == "route" {
// 			routeCookie = v
// 		}
// 	}
// 	if sessionCookie == nil {
// 		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
// 	}
// 	return service.SetUser(service.ZFPrefix, username, password, sessionCookie, routeCookie)
// }

func loginByOauth(username string, password string) (*model.User, error) {
	f := fetch.Fetch{}
	f.Init()
	loginHome, err := f.GetRaw(oauth.OauthLoginHome())
	if err != nil {
		return nil, err
	}
	if len(f.Cookie) < 1 {
		return nil, err
	}
	doc, _ := goquery.NewDocumentFromReader(loginHome.Body)
	hiddenInput := doc.Find("input[type=hidden][name=execution]")

	execution := hiddenInput.AttrOr("value", "")
	loginData := genOauthLoginData(username, password, execution, &f)

	postRedirectUrl, err := f.PostFormRedirect(oauth.OauthLoginHome(), loginData)
	if err != nil {
		return nil, err
	}
	f.Cookie = []*http.Cookie{}
	getRedirectUrl1, err := f.GetRedirect(postRedirectUrl.String())
	if err != nil {
		return nil, err
	}
	getRedirectUrl2, err := f.GetRedirect(getRedirectUrl1.String())
	if err != nil {
		return nil, errors.ERR_OAUTH_NOT_UPDATE
	}
	getRedirectUrl3, err := f.GetRedirect(getRedirectUrl2.String())
	if err != nil {
		return nil, err
	}
	getRedirectUrl4, err := f.GetRedirect(getRedirectUrl3.String())
	if err != nil {
		return nil, err
	}
	_, err = f.Get(getRedirectUrl4.String())
	if err != nil {
		return nil, err
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
