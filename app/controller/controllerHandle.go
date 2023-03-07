package controller

import (
	errors2 "errors"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/utils"
	"funnel/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
		if strings.HasPrefix(context.Request.URL.Path, "/student/zf") {
			if strings.Compare(config.Redis.Get("zf_url").String(), "new") == 0 {
				config.Redis.Set("zf_url", "bk", 0)
			} else {
				config.Redis.Set("zf_url", "new", 0)
			}
		}
	}
	println(err.Error())
	utils.ContextDataResponseJson(context, utils.FailResponseJson(exp, nil))
	context.Abort()
}

func LoginHandle(context *gin.Context, service func(username string, password string, typeFlag bool) (*model.User, error), typeFlag bool) (*model.User, error) {
	user, err := service(context.PostForm("username"), context.PostForm("password"), typeFlag)
	if err != nil {
		ErrorHandle(context, err)
		return nil, err
	}
	return user, err
}
