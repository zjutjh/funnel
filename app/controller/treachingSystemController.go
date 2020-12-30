package controller

import (
	"encoding/json"
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var ts = model.TeachingSystem{}

func ZFLogin(context *gin.Context) {
	session := sessions.Default(context)
	stu := model.ZFUser{
		Username: context.PostForm("username"),
		Password: context.PostForm("password")}

	err := ts.Login(&stu)

	if err == errors.ERR_WRONG_PASSWORD {
		context.Data(200, "application/json", helps.FailResponseJson(errors.WrongPassword, nil))
	} else if err != nil {
		context.Data(200, "application/json", helps.FailResponseJson(errors.UnKnown, nil))
	}

	ZFSession, _ := json.Marshal(stu)
	session.Set("ZFSession", ZFSession)
	_ = session.Save()
	context.Data(200, "application/json", helps.SuccessResponseJson(nil))
}

func GetScoreDetail(context *gin.Context) {
	session := sessions.Default(context)
	ZFSession := session.Get("ZFSession").([]byte)
	if ZFSession == nil {
		context.Data(200, "application/json", helps.FailResponseJson(errors.NotLogin, nil))
		return
	}

	stu := &model.ZFUser{}
	_ = json.Unmarshal(ZFSession, stu)
	result, _ := ts.GetScoreDetail(stu, context.Param("year"), context.Param("term"))
	context.Data(200, "application/json", []byte(result))
}
func GetScore(context *gin.Context) {
	session := sessions.Default(context)
	ZFSession := session.Get("ZFSession").([]byte)
	if ZFSession == nil {
		context.Data(200, "application/json", helps.FailResponseJson(errors.NotLogin, nil))
		return
	}

	stu := &model.ZFUser{}
	_ = json.Unmarshal(ZFSession, stu)
	result, _ := ts.GetScore(stu, context.Param("year"), context.Param("term"))
	context.Data(200, "application/json", []byte(result))
}

func GetClassTable(context *gin.Context) {
	session := sessions.Default(context)
	ZFSession := session.Get("ZFSession").([]byte)
	if ZFSession == nil {
		context.Data(200, "application/json", helps.FailResponseJson(errors.NotLogin, nil))
		return
	}

	stu := &model.ZFUser{}
	_ = json.Unmarshal(ZFSession, stu)
	result, _ := ts.GetClassTable(stu, context.Param("year"), context.Param("term"))
	context.Data(200, "application/json", []byte(result))
}

func GetExamInfo(context *gin.Context) {
	session := sessions.Default(context)
	ZFSession := session.Get("ZFSession").([]byte)
	if ZFSession == nil {
		context.Data(200, "application/json", helps.FailResponseJson(errors.NotLogin, nil))
		return
	}

	stu := &model.ZFUser{}
	_ = json.Unmarshal(ZFSession, stu)
	result, _ := ts.GetExamInfo(stu, context.Param("year"), context.Param("term"))
	context.Data(200, "application/json", []byte(result))
}
