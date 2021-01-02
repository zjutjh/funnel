package controller

import (
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/service"
	"github.com/gin-gonic/gin"
)

var cardSystem = service.CardSystem{}

// @Summary 校园卡余额查询
// @Description 校园卡余额查询
// @Tags 校园卡
// @Produce  json
// @Success 200 json  {"code":200,"data":123.4,"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/card/balance [get]
func CardBalance(context *gin.Context) {
	isValid := helps.CheckPostFormEmpty(
		context,
		[]string{"username", "password"},
	)

	if !isValid {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.InvalidArgs, nil))
		return
	}

	user, err := cardSystem.GetUser(context.PostForm("username"), context.PostForm("password"))
	if err == errors.ERR_WRONG_PASSWORD {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.WrongPassword, nil))
		return
	}
	if err != nil {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
		return
	}

	balance := cardSystem.GetCurrentBalance(user)
	helps.ContextDataResponseJson(context, helps.SuccessResponseJson(balance))
}

// @Summary 校园卡今日消费查询
// @Description 校园卡今日消费查询
// @Tags 校园卡
// @Produce  json
// @Success 200 json  {"code":200,"data":[{}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/card/today [get]
func CardToday(context *gin.Context) {
	isValid := helps.CheckPostFormEmpty(
		context,
		[]string{"username", "password"},
	)

	if !isValid {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.InvalidArgs, nil))
		return
	}

	user, err := cardSystem.GetUser(context.PostForm("username"), context.PostForm("password"))

	if err == errors.ERR_WRONG_PASSWORD {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.WrongPassword, nil))
		return
	}
	if err != nil {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
		return
	}

	balance := cardSystem.GetCardToday(user)
	helps.ContextDataResponseJson(context, helps.SuccessResponseJson(balance))
}

// @Summary 校园卡历史查询
// @Description 校园卡历史查询
// @Tags 校园卡
// @Produce  json
// @Param year body string true "年份"
// @Param month body string true "月份"
// @Success 200 json  {"code":200,"data":[{...}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/card/history [post]
func CardHistory(context *gin.Context) {
	isValid := helps.CheckPostFormEmpty(
		context,
		[]string{"username", "password", "year", "month"},
	)

	if !isValid {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.InvalidArgs, nil))
		return
	}

	user, err := cardSystem.GetUser(context.PostForm("username"), context.PostForm("password"))
	if err == errors.ERR_WRONG_PASSWORD {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.WrongPassword, nil))
		return
	}
	if err != nil {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
		return
	}
	history := cardSystem.GetCardHistory(user, context.PostForm("year"), context.PostForm("month"))
	helps.ContextDataResponseJson(context, helps.SuccessResponseJson(history))
}
