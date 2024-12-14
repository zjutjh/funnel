package libraryService

import (
	"bytes"
	"funnel/app/apis/library"
	"funnel/app/errors"
	"funnel/app/service/libraryService/request"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// OAuthLogin 统一登陆
func OAuthLogin(username string, password string) ([]*http.Cookie, error) {
	client := request.New()
	// 使用cookieJar管理cookie
	cookieJar, _ := cookiejar.New(nil)
	client.SetCookieJar(cookieJar)

	// 1. 初始化请求
	if _, err := client.Request().
		Get(library.OAuthChaoXingInit); err != nil {
		return nil, err
	}

	// 2. 登陆参数生成
	resp, err := client.Request().
		Get(library.LoginFromOAuth)
	if err != nil {
		return nil, err
	}

	// 解析execution
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return nil, err
	}
	execution := doc.
		Find("input[type=hidden][name=execution]").
		AttrOr("value", "")

	// 密码加密
	encPwd, err := GetEncryptedPwd(client, password)

	loginParams := map[string]string{
		"username":   username,
		"mobileCode": "",
		"password":   encPwd,
		"authcode":   "",
		"execution":  execution,
		"_eventId":   "submit",
	}
	// 3. 发送登陆请求
	resp, err = client.Request().
		SetFormData(loginParams).
		Post(library.LoginFromOAuth)
	if err != nil {
		return nil, err
	}

	// 4. 处理重定向
	// 这里我们需要手动的去处理位于js中的重定向
	// resty只能自动处理header.Location中的重定向
	redirect := GetRedirectLocation(resp.String())

	resp, err = client.Request().
		Get(redirect)
	if err != nil {
		return nil, err
	}

	// 5. 提取指定域名下的session并构造cookie列表
	u, _ := url.Parse(library.BaseUrl)
	for _, cookie := range cookieJar.Cookies(u) {
		if cookie.Name == "SESSION" {
			return []*http.Cookie{cookie}, nil
		}
	}
	return nil, errors.ERR_WRONG_PASSWORD
}
