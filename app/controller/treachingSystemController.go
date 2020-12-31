package controller

import (
	"encoding/json"
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/model"
	"funnel/app/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var ts = service.TeachingSystem{}

func ZFLogin(context *gin.Context) {
	_, err := ZFLoginHandle(context)
	if err == errors.ERR_INVALID_ARGS {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.RequestFailed, nil))
	}
	if err == nil {
		helps.ContextDataResponseJson(context, helps.SuccessResponseJson(nil))
	}
}

func GetScoreDetail(context *gin.Context) {

	user, err := ZFTermInfoHandle(context)
	if err == nil {
		result, _ := ts.GetScoreDetail(user, context.PostForm("year"), context.PostForm("term"))
		context.Data(200, "application/json", []byte(result))
	}

}
func GetScore(context *gin.Context) {
	user, err := ZFTermInfoHandle(context)
	if err == nil {
		result, _ := ts.GetScore(user, context.PostForm("year"), context.PostForm("term"))
		context.Data(200, "application/json", []byte(result))
	}
}

func GetClassTable(context *gin.Context) {
	user, err := ZFTermInfoHandle(context)
	if err == nil {
		result, _ := ts.GetClassTable(user, context.PostForm("year"), context.PostForm("term"))
		context.Data(200, "application/json", []byte(result))
	}
}

func GetExamInfo(context *gin.Context) {
	user, err := ZFTermInfoHandle(context)
	if err == nil {
		result, _ := ts.GetExamInfo(user, context.PostForm("year"), context.PostForm("term"))
		context.Data(200, "application/json", []byte(result))
	}
}

func ZFLoginHandle(context *gin.Context) (*model.ZFUser, error) {
	isValid := helps.CheckPostFormEmpty(
		context,
		[]string{"username", "password"},
	)

	if !isValid {
		return nil, errors.ERR_INVALID_ARGS
	}

	user := model.ZFUser{
		Username: context.PostForm("username"),
		Password: context.PostForm("password")}

	err := ts.Login(&user)
	if err == errors.ERR_WRONG_PASSWORD {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.WrongPassword, nil))
		return nil, err
	}
	if err != nil {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
		return nil, err
	}

	session := sessions.Default(context)
	ZFSession, _ := json.Marshal(user)
	session.Set("ZFSession", ZFSession)
	_ = session.Save()
	return &user, nil
}

func ZFTermInfoHandle(context *gin.Context) (*model.ZFUser, error) {
	isValid := helps.CheckPostFormEmpty(
		context,
		[]string{"year", "term"},
	)

	if !isValid {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.RequestFailed, nil))
		return nil, errors.ERR_INVALID_ARGS
	}

	user, err := helps.CheckSession(context, "ZFSession")

	if err != nil {
		user, loginErr := ZFLoginHandle(context)
		if loginErr == nil {
			return user, nil
		}
		if err == errors.ERR_SESSION_NOT_EXIST {
			helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.NotLogin, nil))
		}
		if err == errors.ERR_SESSION_EXPIRES {
			helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.NotLogin, nil))
		}
		return nil, err

	}
	return user, nil
}
