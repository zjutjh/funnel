package zfClient

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"funnel/captcha"
	"funnel/comm"
	"image"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/zjutjh/mygo/nesty"
	"github.com/zjutjh/mygo/nlog"
)

const prefix = "biz.zf"

type ZFClient struct {
	HTTPClient *resty.Client
	Logger     *logrus.Logger
}

func New(ctx context.Context) *ZFClient {
	return &ZFClient{
		HTTPClient: nesty.Pick(prefix).SetBaseURL(comm.BizConf.ZF.BaseURL),
		Logger:     nlog.Pick().WithContext(ctx).Logger,
	}
}

func (c *ZFClient) R() *resty.Request {
	return c.HTTPClient.R()
}

// BypassCaptcha 获取成功识别验证码的cookie, 登录使用
func (c *ZFClient) BypassCaptcha() ([]*http.Cookie, error) {
	// 初始化登录
	response, err := c.R().SetQueryParam("type", "resource").Get(comm.ZFLoginCaptchaLogin)
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
	var captchaData CaptchaLoginResponse
	response, err = c.R().
		SetCookies(cookies).
		SetQueryParam("type", "refresh").
		SetResult(&captchaData).
		Get(comm.ZFCaptchaLogin)
	if err != nil {
		c.Logger.Errorf("正方获取验证码失败: %v", err)
		return nil, err
	}

	imgResp, err := c.R().
		SetCookies(cookies).
		SetQueryParams(map[string]string{
			"type": "image",
			"id":   captchaData.Si,
			"imtk": captchaData.Imtk,
		}).
		Get(comm.ZFCaptchaLogin)
	if err != nil {
		c.Logger.Errorf("正方下载验证码图片失败: %v", err)
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(imgResp.Body()))
	if err != nil {
		c.Logger.Errorf("正方下载验证码图片失败: %v", err)
		return nil, err
	}
	result, err := captcha.Crack(img)
	if err != nil {
		c.Logger.Errorf("正方验证码识别失败: %v", err)
		return nil, err
	}
	var verifyData CaptchaVerifyResponse
	response, err = c.R().
		SetCookies(cookies).
		SetResult(&verifyData).
		SetQueryParams(map[string]string{
			"type":   "verify",
			"rtk":    rtk,
			"mt":     base64.StdEncoding.EncodeToString([]byte(result)),
			"extend": "eyJhcHBOYW1lIjoiTmV0c2NhcGUiLCJ1c2VyQWdlbnQiOiJNb3ppbGxhLzUuMCAoTWFjaW50b3NoOyBJbnRlbCBNYWMgT1MgWCAxMF8xNV83KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWUvMTQ0LjAuMC4wIFNhZmFyaS81MzcuMzYiLCJhcHBWZXJzaW9uIjoiNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xNDQuMC4wLjAgU2FmYXJpLzUzNy4zNiJ9",
		}).Post(comm.ZFCaptchaLogin)
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
	return cookies, err
}
