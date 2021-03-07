package controller

import (
	errors2 "errors"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
			break
		}
	case errors.ERR_WRONG_Captcha:
		{
			exp = errors.CaptchaFailed
			break
		}
	case errors.ERR_Session_Expired:
		{
			exp = errors.SessionExpired
			break
		}

	}
	if errors2.Is(err, http.ErrHandlerTimeout) {
		exp = errors.RequestFailed
	}

	utils.ContextDataResponseJson(context, utils.FailResponseJson(exp, nil))
	context.Abort()
}

func LoginHandle(context *gin.Context, service func(username string, password string) (*model.User, error)) (*model.User, error) {
	user, err := service(context.PostForm("username"), context.PostForm("password"))
	if err != nil {
		ErrorHandle(context, err)
		return nil, err
	}
	return user, err
}
