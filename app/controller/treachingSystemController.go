package controller

import (
	"encoding/json"
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/model"
	"funnel/app/service"
	"github.com/gin-gonic/gin"
	"log"
)

var ts = service.TeachingSystem{}

// @Summary 正方教务详细成绩
// @Description 正方教务详细成绩
// @Tags 正方
// @Produce  json
// @Param term body string true "学期"
// @Param year body string true "年份"
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":{object},"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/zf/score/info [post]
func GetScoreDetail(context *gin.Context) {
	user, err := ZFTermInfoHandle(context)
	if err == nil {
		result, err2 := ts.GetScoreDetail(user, context.PostForm("year"), context.PostForm("term"))
		if err2 == nil {
			var f interface{}
			_ = json.Unmarshal([]byte(result), &f)
			helps.ContextDataResponseJson(context, helps.SuccessResponseJson(f))
			return
		}
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
	}
}

// @Summary 正方教务成绩
// @Description 正方教务成绩
// @Tags 正方
// @Produce  json
// @Param term body string true "学期"
// @Param year body string true "年份"
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":{object},"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/zf/score/info [post]
func GetScore(context *gin.Context) {
	user, err := ZFTermInfoHandle(context)
	if err == nil {
		result, err := ts.GetScore(user, context.PostForm("year"), context.PostForm("term"))
		if err == nil {
			var f interface{}
			_ = json.Unmarshal([]byte(result), &f)
			helps.ContextDataResponseJson(context, helps.SuccessResponseJson(f))
			return
		}
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
	}
}

// @Summary 正方教务课表
// @Description 正方教务课表
// @Tags 正方
// @Produce  json
// @Param term body string true "学期"
// @Param year body string true "年份"
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":{object},"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/zf/score [post]
func GetClassTable(context *gin.Context) {
	user, err := ZFTermInfoHandle(context)
	if err == nil {
		result, err := ts.GetClassTable(user, context.PostForm("year"), context.PostForm("term"))
		if err == nil {
			var f interface{}
			_ = json.Unmarshal([]byte(result), &f)
			helps.ContextDataResponseJson(context, helps.SuccessResponseJson(f))
			return
		}
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
	}
}

// @Summary 正方教务考试信息
// @Description 正方教务考试信息
// @Tags 正方
// @Produce  json
// @Param term body string true "学期"
// @Param year body string true "年份"
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":{object},"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/zf/exam [post]
func GetExamInfo(context *gin.Context) {
	user, err := ZFTermInfoHandle(context)
	if err == nil {
		result, err := ts.GetExamInfo(user, context.PostForm("year"), context.PostForm("term"))
		if err == nil {
			var f interface{}
			_ = json.Unmarshal([]byte(result), &f)
			helps.ContextDataResponseJson(context, helps.SuccessResponseJson(f))
			return
		}
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
	}
}

// @Summary 正方教务考试信息
// @Description 正方教务考试信息
// @Tags 正方
// @Produce  json
// @Param term body string true "学期"
// @Param year body string true "年份"
// @Param campus body string true "校区"
// @Param weekday body string true "星期几 1，2，3"
// @Param week body string true "第几周的2次幂的和"
// @Param classPeriod body string true "第几节课的2次幂的和"
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":{object},"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/zf/exam [post]
func GetRoomInfo(context *gin.Context) {
	user, err := ZFTermInfoHandle(context)
	if err == nil {
		isValid := helps.CheckPostFormEmpty(
			context,
			[]string{"week", "classPeriod", "campus", "weekday"},
		)

		if !isValid {
			helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.InvalidArgs, nil))
			return
		}

		result, err := ts.GetEmptyRoomInfo(
			context.PostForm("year"),
			context.PostForm("term"),
			context.PostForm("campus"),
			context.PostForm("weekday"),
			context.PostForm("week"),
			context.PostForm("classPeriod"),
			user,
		)
		if err == nil {
			var f interface{}
			_ = json.Unmarshal([]byte(result), &f)
			helps.ContextDataResponseJson(context, helps.SuccessResponseJson(f))
			return
		}
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
	}
}

func ZFTermInfoHandle(context *gin.Context) (*model.User, error) {
	isValid := helps.CheckPostFormEmpty(
		context,
		[]string{"username", "password", "year", "term"},
	)

	if !isValid {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.InvalidArgs, nil))
		return nil, errors.ERR_INVALID_ARGS
	}
	user, err := ts.GetUser(context.PostForm("username"), context.PostForm("password"))
	if err == errors.ERR_WRONG_PASSWORD {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.WrongPassword, nil))
		return nil, err
	}
	if err == errors.ERR_WRONG_Captcha {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.CaptchaFailed, nil))
		return nil, err
	}
	if err != nil {
		log.Println(err.Error())
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
		return nil, err
	}
	return user, nil
}
