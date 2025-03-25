package zfService

import (
	"funnel/app/errors"
	"regexp"

	"github.com/go-resty/resty/v2"
)

// 登陆错误对应在页面的提示信息
const (
	WrongPasswordMsg = "用户名或密码错误" // #nosec G101
	WrongAccountMsg  = "当前账号无权登录"
	NotActivatedMsg  = "账号未激活，请激活后再登录"
)

// GetLoginMsg 获取登陆失败后页面上的提示语
func GetLoginMsg(resp *resty.Response) string {
	re := regexp.MustCompile(`<span\s+id="msg">(.+?)</span>`)
	matches := re.FindStringSubmatch(resp.String())
	if len(matches) == 0 {
		return ""
	}
	// 删除span内部的标签
	re = regexp.MustCompile(`<[^>]*>`)
	msg := re.ReplaceAllString(matches[1], "")
	return msg
}

// CheckLogin 用于判断登陆是否成功
func CheckLogin(resp *resty.Response) error {
	// 判断失败原因
	msg := GetLoginMsg(resp)
	if msg == "" {
		return nil
	}
	switch msg {
	case WrongPasswordMsg:
		return errors.ERR_WRONG_PASSWORD
	case WrongAccountMsg:
		return errors.ERR_WRONG_PASSWORD
	case NotActivatedMsg:
		return errors.ERR_OAUTH_NOT_UPDATE
	}
	return errors.ERR_UNKNOWN_LOGIN_ERROR
}
