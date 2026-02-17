package zfClient

// TODO  SessionPool

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"funnel/captcha"
	"funnel/comm"
	"image"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/zjutjh/mygo/nlog"
)

type ZFClient struct {
	Logger *logrus.Logger
}

var (
	httpClient *resty.Client
	once       sync.Once
)

func New(ctx context.Context) *ZFClient {
	once.Do(func() {
		httpClient = resty.New()
		httpClient.
			SetBaseURL(comm.BizConf.ZF.BaseURL).
			SetRedirectPolicy(resty.NoRedirectPolicy())
	})
	return &ZFClient{
		Logger: nlog.Pick().WithContext(ctx).Logger,
	}
}

func (c *ZFClient) R() *resty.Request {
	return httpClient.R()
}

// BypassCaptcha 获取成功识别验证码的cookie, 登录使用
// 1. 获取登录初始化信息，包含rtk和初始cookie
// 2. 获取验证码信息，包含si和imtk
// 3. 下载验证码图片并识别
// 4. 验证验证码结果，成功则返回cookie，失败则返回错误
func (c *ZFClient) BypassCaptcha() (*ZFCookie, error) {
	// 初始化登录
	response, err := c.R().
		SetQueryParams(map[string]string{
			"type": "resource",
			"name": "zfdun_captcha.js",
		}).
		Get(comm.ZFCaptchaURL)
	if err != nil {
		c.Logger.Errorf("正方验证码登陆初始化失败: %v", err)
		return nil, err
	}
	text := response.String()
	rtk := comm.ExtractJsonStr(text, "rtk")
	if rtk == "" {
		err = fmt.Errorf("rtk解析失败")
		c.Logger.Errorf("正方验证码登陆初始化失败: %v", err)
		return nil, err
	}
	cookies := response.Cookies()

	// 获取验证码
	var captchaData captchaLoginResponse
	response, err = c.R().
		SetCookies(cookies).
		SetQueryParam("type", "refresh").
		SetResult(&captchaData).
		Get(comm.ZFCaptchaURL)
	if err != nil {
		c.Logger.Errorf("正方获取验证码失败: %v", err)
		return nil, err
	}

	// 下载验证码图片
	imgResp, err := c.R().
		SetCookies(cookies).
		SetQueryParams(map[string]string{
			"type": "image",
			"id":   captchaData.Si,
			"imtk": captchaData.Imtk,
		}).
		Get(comm.ZFCaptchaURL)
	if err != nil {
		c.Logger.Errorf("正方下载验证码图片失败: %v", err)
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(imgResp.Body()))
	if err != nil {
		c.Logger.Errorf("正方下载验证码图片失败: %v", err)
		return nil, err
	}

	// 识别
	result, err := captcha.Crack(img)
	if err != nil {
		c.Logger.Errorf("正方验证码识别失败: %v", err)
		return nil, err
	}
	var verifyData captchaVerifyResponse
	response, err = c.R().
		SetCookies(cookies).
		SetResult(&verifyData).
		SetQueryParams(map[string]string{
			"type":   "verify",
			"rtk":    rtk,
			"mt":     base64.StdEncoding.EncodeToString([]byte(result)),
			"extend": "eyJhcHBOYW1lIjoiTmV0c2NhcGUiLCJ1c2VyQWdlbnQiOiJNb3ppbGxhLzUuMCAoTWFjaW50b3NoOyBJbnRlbCBNYWMgT1MgWCAxMF8xNV83KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWUvMTQ0LjAuMC4wIFNhZmFyaS81MzcuMzYiLCJhcHBWZXJzaW9uIjoiNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xNDQuMC4wLjAgU2FmYXJpLzUzNy4zNiJ9",
		}).Post(comm.ZFCaptchaURL)
	if err != nil {
		c.Logger.Errorf("正方验证码验证失败: %v", err)
		return nil, err
	}
	success := verifyData.Status == "success"
	if !success {
		err = fmt.Errorf("验证码验证失败: %s", verifyData.Msg)
		c.Logger.Errorf("正方验证码验证失败: %v", err)
		return nil, err
	}
	c.Logger.Infof("验证码验证结果: %s", response.String())
	return FromCookie(cookies), err
}

// LoginByCaptcha 验证码登录正方
// 1. 获取成功识别验证码的cookie
// 2. 加密密码
// 3. 提交登录请求，成功则返回cookie，失败则返回错误
func (c *ZFClient) LoginByCaptcha(username, password string) (*ZFCookie, error) {
	zfCookies, err := c.BypassCaptcha()
	if err != nil {
		return nil, err
	}
	cookies := zfCookies.ToCookie()
	var publicKey LoginPublicKeyResponse
	_, err = c.R().SetCookies(cookies).
		SetResult(&publicKey).
		Get(comm.ZFLoginPublicKeyURL)
	if err != nil {
		c.Logger.Errorf("正方获取登录公钥失败: %v", err)
		return nil, err
	}
	encryptedPassword, err := comm.RSAEncryptWithPublicKey(password, publicKey.Modulus, publicKey.Exponent)
	if err != nil {
		c.Logger.Errorf("正方加密密码失败: %v", err)
		return nil, err
	}
	response, err := c.R().SetCookies(cookies).
		SetQueryParams(map[string]string{
			"yhm": username,
			"mm":  encryptedPassword,
		}).
		Post(comm.ZFLoginURL)
	if err != nil && !errors.Is(err, resty.ErrAutoRedirectDisabled) {
		return nil, err
	}
	if response.StatusCode() != http.StatusFound {
		// 登录失败
		return nil, checkError(response.String())
	}

	// merge cookie
	zfCookie := FromCookie(cookies)
	for _, cookie := range response.Cookies() {
		if cookie.Name == "JSESSIONID" {
			zfCookie.JSessionID = cookie.Value
		}
	}
	// 登录成功
	return zfCookie, nil
}

// GetCurrentSchoolTerm 查询正方当前学年学期信息
func (c *ZFClient) GetCurrentSchoolTerm(cookie *ZFCookie) (*SchoolTermInfo, error) {
	response, err := c.R().SetCookies(cookie.ToCookie()).Get(comm.ZFLessonTableHome)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body()))
	if err != nil {
		return nil, err
	}
	Year := doc.Find("#xnm option[selected]").AttrOr("value", "")
	Term := doc.Find("#xqm option[selected]").AttrOr("value", "")
	return &SchoolTermInfo{
		Year: Year,
		Term: Term,
	}, nil
}

func checkError(html string) error {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return fmt.Errorf("解析 HTML 失败: %v", err)
	}
	text := strings.TrimSpace(doc.Find("#tips").Text())
	switch text {
	case "用户名或密码不正确，请重新输入！":
		return comm.WrongUsernameOrPasswordError
	default:
		return comm.UnknownError
	}
}
