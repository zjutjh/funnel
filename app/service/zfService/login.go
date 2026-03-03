package zfService

import (
	"bytes"
	"encoding/base64"
	stderrors "errors"
	"fmt"
	"funnel/app/apis/oauth"
	"funnel/app/apis/zf"
	"funnel/app/captcha"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils/fetch"
	"funnel/app/utils/security"
	"image"
	"log/slog"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

type captchaLoginResponse struct {
	Msg    string `json:"msg"`
	T      int64  `json:"t"`
	Si     string `json:"si"`
	Imtk   string `json:"imtk"`
	Mi     string `json:"mi"`
	Vs     string `json:"vs"`
	Status string `json:"status"`
}

type captchaVerifyResponse struct {
	Msg    string `json:"msg"`
	Vs     string `json:"vs"`
	Status string `json:"status"`
}

type LoginPublicKeyResponse struct {
	Modulus  string `json:"modulus"`
	Exponent string `json:"exponent"`
}

func extractJsonStr(jsonStr, jsonKey string) string {
	re := regexp.MustCompile(jsonKey + `:'([^']+)'`)
	matches := re.FindStringSubmatch(jsonStr)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// BypassCaptcha 获取成功识别验证码的cookie, 登录使用
// 1. 获取登录初始化信息，包含rtk和初始cookie
// 2. 获取验证码信息，包含si和imtk
// 3. 下载验证码图片并识别
// 4. 验证验证码结果，成功则返回cookie，失败则返回错误
func bypassCaptcha() ([]*http.Cookie, error) {
	client := resty.New().SetCookieJar(nil)
	// 初始化登录
	response, err := client.R().
		SetQueryParams(map[string]string{
			"type": "resource",
			"name": "zfdun_captcha.js",
		}).
		Get(zf.ZfCaptchaURL())
	if err != nil {
		slog.Error("正方验证码登陆初始化失败", "err", err)
		return nil, err
	}
	text := response.String()
	rtk := extractJsonStr(text, "rtk")
	if rtk == "" {
		err = fmt.Errorf("rtk解析失败")
		slog.Error("正方验证码登陆初始化失败", "err", err)
		return nil, err
	}
	cookies := response.Cookies()

	// 获取验证码
	var captchaData captchaLoginResponse
	response, err = client.R().
		SetCookies(cookies).
		SetQueryParam("type", "refresh").
		SetResult(&captchaData).
		Get(zf.ZfCaptchaURL())
	if err != nil {
		slog.Error("正方获取验证码失败", "err", err)
		return nil, err
	}

	// 下载验证码图片
	imgResp, err := client.R().
		SetCookies(cookies).
		SetQueryParams(map[string]string{
			"type": "image",
			"id":   captchaData.Si,
			"imtk": captchaData.Imtk,
		}).
		Get(zf.ZfCaptchaURL())
	if err != nil {
		slog.Error("正方下载验证码图片失败", "err", err)
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(imgResp.Body()))
	if err != nil {
		slog.Error("正方下载验证码图片失败", "err", err)
		return nil, err
	}

	// 识别
	result, err := captcha.Crack(img)
	if err != nil {
		slog.Error("正方验证码识别失败", "err", err)
		return nil, err
	}
	var verifyData captchaVerifyResponse
	response, err = client.R().
		SetCookies(cookies).
		SetResult(&verifyData).
		SetQueryParams(map[string]string{
			"type":   "verify",
			"rtk":    rtk,
			"mt":     base64.StdEncoding.EncodeToString([]byte(result)),
			"extend": "eyJhcHBOYW1lIjoiTmV0c2NhcGUiLCJ1c2VyQWdlbnQiOiJNb3ppbGxhLzUuMCAoTWFjaW50b3NoOyBJbnRlbCBNYWMgT1MgWCAxMF8xNV83KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWUvMTQ0LjAuMC4wIFNhZmFyaS81MzcuMzYiLCJhcHBWZXJzaW9uIjoiNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xNDQuMC4wLjAgU2FmYXJpLzUzNy4zNiJ9",
		}).Post(zf.ZfCaptchaURL())
	if err != nil {
		slog.Error("正方验证码验证失败", "err", err)
		return nil, err
	}
	success := verifyData.Status == "success"
	if !success {
		err = fmt.Errorf("验证码验证失败: %s", verifyData.Msg)
		slog.Error("正方验证码验证失败", "err", err)
		return nil, err
	}
	slog.Info("验证码验证结果:", "result", response.String())
	return cookies, nil
}

// login 验证码登录正方
// 1. 获取成功识别验证码的cookie
// 2. 加密密码
// 3. 提交登录请求，成功则返回用户信息，失败则返回错误
func login(username, password string) (*model.User, error) {
	cookies, err := bypassCaptcha()
	if err != nil {
		return nil, err
	}
	client := resty.New()
	client.SetRedirectPolicy(resty.NoRedirectPolicy())
	client.SetCookieJar(nil)

	// 获取公钥
	var publicKey LoginPublicKeyResponse
	_, err = client.R().SetCookies(cookies).
		SetResult(&publicKey).
		Get(zf.ZfLoginGetPublickey())
	if err != nil {
		slog.Error("正方获取登录公钥失败", "err", err)
		return nil, err
	}
	// 加密密码
	encryptedPassword, err := security.RSAEncryptWithPublicKey(password, publicKey.Modulus, publicKey.Exponent)
	if err != nil {
		slog.Error("正方加密密码失败", "err", err)
		return nil, err
	}
	// 提交登录
	response, err := client.R().
		SetCookies(cookies).
		SetFormData(map[string]string{
			"yhm": username,
			"mm":  encryptedPassword,
		}).
		Post(zf.ZfLoginHome())
	if err != nil && !stderrors.Is(err, resty.ErrAutoRedirectDisabled) {
		return nil, err
	}
	if response.StatusCode() != http.StatusFound {
		s := response.Body()
		if strings.Contains(string(s), "请先滑动图片进行验证") {
			return nil, errors.ERR_WRONG_Captcha
		}
		if strings.Contains(string(s), "用户名或密码不正确") {
			return nil, errors.ERR_WRONG_PASSWORD
		}
		// 登录失败
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}
	// 提取cookie
	var sessionCookie *http.Cookie
	var routeCookie *http.Cookie
	for _, cookie := range response.Cookies() {
		if cookie.Name == "JSESSIONID" {
			sessionCookie = cookie
		}
	}

	for _, cookie := range cookies {
		if cookie.Name == "route" {
			routeCookie = cookie
		}
	}
	if sessionCookie == nil || routeCookie == nil {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}
	return service.SetUser(service.ZFPrefix, username, password, sessionCookie, routeCookie)
}

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
