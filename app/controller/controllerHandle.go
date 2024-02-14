package controller

import (
	"fmt"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/utils"
	"funnel/config/logs"
	"github.com/gin-gonic/gin"
)

func ErrorHandle(context *gin.Context, err error) {
	if err == nil {
		return
	}
	var exp = errors.UnKnown
	switch err {
	case errors.ERR_WRONG_PASSWORD:
		{
			exp = errors.WrongPassword
			logs.WriteDebug(context, exp)
			break
		}
	case errors.ERR_WRONG_Captcha:
		{
			exp = errors.CaptchaFailed
			logs.WriteWarn(context, exp)
			break
		}
	case errors.ERR_Session_Expired:
		{
			exp = errors.SessionExpired
			logs.WriteWarn(context, exp)
			break
		}
	default:
		{
			logs.WriteError(context, err)
		}
	}

	println(err.Error())
	utils.ContextDataResponseJson(context, utils.FailResponseJson(exp, nil))
	context.Abort()
}

func LoginHandle(context *gin.Context, service func(username, password string, loginType LoginType, typeFlag bool) (*model.User, error), typeFlag bool) (*model.User, error) {
	loginType, err := ParseLoginType(context.PostForm("type"))
	if err != nil {
		ErrorHandle(context, err)
		return nil, err
	}
	user, err := service(context.PostForm("username"), context.PostForm("password"), loginType, typeFlag)
	if err != nil {
		ErrorHandle(context, err)
		return nil, err
	}
	return user, err
}

var loginTypeNames = map[string]LoginType{
	"ZF":    ZF,
	"OAUTH": OAUTH,
}

func ParseLoginType(str string) (LoginType, error) {
	status, found := loginTypeNames[str]
	if !found {
		return -1, fmt.Errorf("invalid status: %s", str)
	}
	return status, nil
}

type LoginType int

const (
	ZF LoginType = iota
	OAUTH
)
