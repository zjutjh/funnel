package controller

import (
	"encoding/json"
	funnel "funnel/app"
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var ts = funnel.TeachingSystem{RootUrl: "http://172.16.19.163/jwglxt/"}

func ZFLogin(context *gin.Context) {
	session := sessions.Default(context)
	stu := model.Student{Sid: context.PostForm("username"), Password: context.PostForm("password")}

	isLoginSuccess := false
	for i := 0; i < 10; i++ {
		err := ts.Login(&stu)
		if err == errors.ERR_WRONG_PASSWORD {
			context.Data(200, "application/json", helps.FailResponseJson(errors.WrongPassword, nil))
		}
		if err == nil {
			isLoginSuccess = true
			break
		}

	}

	if isLoginSuccess == false {
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
		context.Data(200, "application/json", []byte("error"))
		return
	}

	stu := &model.Student{}
	_ = json.Unmarshal(ZFSession, stu)
	context.Data(200, "application/json", []byte(ts.GetScoreDetail(stu, context.Param("year"), context.Param("term"))))
}
func GetScore(context *gin.Context) {
	session := sessions.Default(context)
	ZFSession := session.Get("ZFSession").([]byte)
	if ZFSession == nil {
		context.Data(200, "application/json", helps.FailResponseJson(errors.NotLogin, nil))
		return
	}

	stu := &model.Student{}
	_ = json.Unmarshal(ZFSession, stu)
	context.Data(200, "application/json", []byte(ts.GetScore(stu, context.Param("year"), context.Param("term"))))
}

func GetClassTable(context *gin.Context) {
	session := sessions.Default(context)
	ZFSession := session.Get("ZFSession").([]byte)
	if ZFSession == nil {
		context.Data(200, "application/json", []byte("error"))
		return
	}

	stu := &model.Student{}
	_ = json.Unmarshal(ZFSession, stu)
	context.Data(200, "application/json", []byte(ts.GetClassTable(stu, context.Param("year"), context.Param("term"))))
}

func GetExamInfo(context *gin.Context) {
	session := sessions.Default(context)
	ZFSession := session.Get("ZFSession").([]byte)
	if ZFSession == nil {
		context.Data(200, "application/json", []byte("error"))
		return
	}

	stu := &model.Student{}
	_ = json.Unmarshal(ZFSession, stu)
	context.Data(200, "application/json", []byte(ts.GetExamInfo(stu, context.Param("year"), context.Param("term"))))
}
