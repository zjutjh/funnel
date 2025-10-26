package zfLoginTokenRegister

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type ZFLoginTokenRegister struct {
	CookieJSESSIONID string
	CookieRoute      string
	CSRFToken        string // 登录所需的 CSRFToken
	CryptoModulus    string // 登录所需的公钥m值
	CryptoExponent   string // 登录所需的公钥e值
	hostUrl          string // 主机地址 外部传入
	instanceId       string // 验证码实例ID 固定值zfcaptchaLogin
	userAgent        string // 模拟的User-Agent字段 外部传入
	appVersionUA     string // 模拟的appVersion字段 目前和UA相同 非程序版本
	restyClient      *resty.Client
}

type captchaImgTokens struct {
	Si   string // 背景图ID
	Mi   string // 滑块图ID
	Imtk string // 图片请求令牌
}

// PUBLIC METHOD

func (r *ZFLoginTokenRegister) Init(hostUrl string, ua string) {
	r.hostUrl = strings.TrimSuffix(hostUrl, "/")
	r.userAgent = ua
	r.appVersionUA = ua
	r.instanceId = "zfcaptchaLogin" // 抓包获取的固定值 写死的常量
	r.restyClient = resty.New()
	r.restyClient.SetHeader("User-Agent", ua)
	r.restyClient.SetCookieJar(nil)
}

// RunAndWaitCaptcha 初始化登录请求并拿到未通过验证码的 token
// 返回验证码图片供外部操作
// 这是一个耗时操作
func (r *ZFLoginTokenRegister) RunAndWaitCaptchaBg() (image.Image, error) {
	// 请求登录页面 获取初始会话
	if err := r.getInitialSession(); err != nil {
		return nil, err
	}
	// 获取验证码图片的相关令牌
	imgToken, err := r.getCaptchaImgTokens()
	if err != nil {
		return nil, err
	}
	// 下载验证码背景图
	if img, err := r.getCaptchaImgBg(imgToken.Si, imgToken.Imtk); err != nil {
		return nil, err
	} else {
		return img, nil
	}
}

// SolveAndSubmit 提交验证码结果 如果请求正常并通过验证则返回 true
// 参数 captchaSolution 是破解器返回的轨迹包字符串
// 这是一个耗时操作
func (r *ZFLoginTokenRegister) SolveAndSubmit(captchaSolution string) (bool, error) {
	// 获取 rtk 令牌
	rtk, err := r.GetRtkToken()
	if err != nil {
		return false, err
	}
	// 提交验证码结果
	isPassed, err := r.submitCaptcha(rtk, captchaSolution)
	if err != nil {
		return false, err
	}
	if !isPassed {
		return false, fmt.Errorf("captcha not passed") // TODO RETRY
	}
	return true, nil
}

// PRIVATE METHOD

// getInitialSession 访问登录页以获取初始的 JSESSIONID，这是后续所有请求的基础。
func (r *ZFLoginTokenRegister) getInitialSession() error {
	// 构造请求
	reqURL := r.hostUrl + "/jwglxt/xtgl/s_login.html"
	resp, err := r.restyClient.R().Get(reqURL)
	if err != nil {
		return fmt.Errorf("try get login page error: %w", err)
	}
	// 手动获取并储存 cookie
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "JSESSIONID" {
			r.CookieJSESSIONID = cookie.Value
		}
		if cookie.Name == "route" {
			r.CookieRoute = cookie.Value
		}
	}
	// 检查是否成功获取 cookie
	if r.CookieJSESSIONID == "" || r.CookieRoute == "" {
		return fmt.Errorf("cannot get JSESSIONID or route from login page")
	}
	return nil
}

// getCaptchaImgTokens 请求用于获取验证码图片的相关令牌
func (r *ZFLoginTokenRegister) getCaptchaImgTokens() (captchaImgTokens, error) {
	type imageTokens struct { // 定义内部映射结构体
		Si     string `json:"si"`     // 背景图ID
		Mi     string `json:"mi"`     // 滑块图ID (未使用)
		Imtk   string `json:"imtk"`   // 图片请求令牌
		T      int64  `json:"t"`      // 时间戳
		Status string `json:"status"` // 返回的业务状态
		Msg    string `json:"msg"`    // 返回的业务消息
	}
	var imgTokens imageTokens
	cookie := fmt.Sprintf("JSESSIONID=%s; route=%s", r.CookieJSESSIONID, r.CookieRoute)
	resp, err := r.restyClient.R().
		SetHeader("Cookie", cookie).
		SetResult(&imgTokens).
		Get(r.hostUrl + "/jwglxt/zfcaptchaLogin?instanceId=zfcaptchaLogin&type=refresh")
	if err != nil || resp.IsError() {
		return captchaImgTokens{}, fmt.Errorf("try get img tokens failed: %w", err)
	}
	if imgTokens.Status != "success" {
		return captchaImgTokens{}, fmt.Errorf("err happened when getting img token: %s", imgTokens.Msg)
	}
	return captchaImgTokens{
		Si:   imgTokens.Si,
		Mi:   imgTokens.Mi,
		Imtk: imgTokens.Imtk,
	}, nil
}

// getCaptchaImgBg 下载验证码背景图 实际是 downloadCaptchaImg 的包装
func (r *ZFLoginTokenRegister) getCaptchaImgBg(si, imtk string) (image.Image, error) {
	return r.downloadCaptchaImg(si, imtk, strconv.FormatInt(time.Now().UnixMicro(), 10))
}

// downloadCaptchaImg 根据图片ID等参数下载验证码图片
func (r *ZFLoginTokenRegister) downloadCaptchaImg(imgId, imtk, t string) (image.Image, error) {
	query := map[string]string{
		"type":       "image",
		"id":         imgId,
		"imtk":       imtk,
		"t":          t,
		"instanceId": r.instanceId,
	}

	// 构造请求并发送
	cookie := fmt.Sprintf("JSESSIONID=%s; route=%s", r.CookieJSESSIONID, r.CookieRoute)
	resp, err := r.restyClient.R().
		SetQueryParams(query).
		SetHeader("Cookie", cookie).
		Get(r.hostUrl + "/jwglxt/zfcaptchaLogin")
	if err != nil || resp.IsError() {
		return nil, fmt.Errorf("download captcha img (id=%s) failed: %w", imgId, err)
	}

	// 解码图片并返回
	if img, _, err := image.Decode(bytes.NewReader(resp.Body())); err != nil {
		return nil, fmt.Errorf("captcha image decode (id=%s) failed: %w", imgId, err)
	} else {
		return img, nil
	}
}

// GetRtkToken 下载包含 rtk 令牌的 JS 文件并提取 rtk
func (r *ZFLoginTokenRegister) GetRtkToken() (string, error) {
	cookie := fmt.Sprintf("JSESSIONID=%s; route=%s", r.CookieJSESSIONID, r.CookieRoute)
	resp, err := r.restyClient.R().
		SetHeader("Cookie", cookie).
		Get(r.hostUrl + "/jwglxt/zfcaptchaLogin?type=resource&instanceId=zfcaptchaLogin&name=zfdun_captcha.js")
	if err != nil || resp.IsError() {
		return "", fmt.Errorf("cannot download rtk-included js: %w", err)
	}
	// 只取前256字节以提高效率
	responseText := resp.Body()
	responseText = responseText[0:256]
	// 正则提取 rtk
	re := regexp.MustCompile(`rtk:'([a-f0-9\-]+)'`)
	matches := re.FindStringSubmatch(string(responseText))
	// 检查是否找到匹配项 由于实际 rtk 在捕获组中 因此至少应有两个元素
	if len(matches) < 2 {
		return "", fmt.Errorf("在 JS 文件中未找到 RTK 令牌")
	}
	// 由于正则捕获组从索引 1 开始，因此返回 matches[1]
	return matches[1], nil
}

// submitCaptcha 提交破解得到的验证码轨迹包
func (r *ZFLoginTokenRegister) submitCaptcha(rtk, mt string) (bool, error) {
	// 构造 extend mt 的 base64 值
	extendData := map[string]string{
		"appName":    "Netscape",
		"userAgent":  r.userAgent,
		"appVersion": r.appVersionUA,
	}
	extendBytes, _ := json.Marshal(extendData)
	extendB64 := base64.StdEncoding.EncodeToString(extendBytes)
	mtB64 := base64.StdEncoding.EncodeToString([]byte(mt))

	// 组装表单数据
	formData := map[string]string{
		"type":       "verify",
		"rtk":        rtk,
		"time":       strconv.FormatInt(time.Now().UnixMilli(), 10),
		"mt":         mtB64,
		"instanceId": r.instanceId,
		"extend":     extendB64,
	}
	cookie := fmt.Sprintf("JSESSIONID=%s; route=%s", r.CookieJSESSIONID, r.CookieRoute)

	var verifyRes struct { // 定义响应映射临时结构体
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	resp, err := resty.New().R().
		SetFormData(formData).
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetHeader("Cookie", cookie).
		SetResult(&verifyRes).
		Post(r.hostUrl + "/jwglxt/zfcaptchaLogin")
	if err != nil || resp.IsError() {
		return false, fmt.Errorf("cannot submit captcha: %w", err)
	}

	// 业务状态码判断
	if verifyRes.Status != "success" {
		return false, fmt.Errorf("captcha not passed: %s", verifyRes.Message)
	}
	return true, nil
}
