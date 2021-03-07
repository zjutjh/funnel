package schoolCardController

import (
	"funnel/app/controller"
	"funnel/app/errors"
	"funnel/app/service/schoolcardService"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 校园卡余额查询
// @Description 校园卡余额查询
// @Tags 校园卡
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":123.4,"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/card/balance [post]
func CardBalance(context *gin.Context) {
	user, err := controller.LoginHandle(context, schoolcardService.GetUser)
	if err != nil {
		return
	}
	balance, err := schoolcardService.GetCurrentBalance(user)
	if err == errors.ERR_Session_Expired {
		user, err = controller.LoginHandle(context, schoolcardService.GetUser)
		balance, err = schoolcardService.GetCurrentBalance(user)
	}

	if err != nil {
		controller.ErrorHandle(context, err)
		return
	}

	utils.ContextDataResponseJson(context, utils.SuccessResponseJson(balance))
}

// @Summary 校园卡今日消费查询
// @Description 校园卡今日消费查询
// @Tags 校园卡
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":[{}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/card/today [post]
func CardToday(context *gin.Context) {
	user, err := controller.LoginHandle(context, schoolcardService.GetUser)
	if err != nil {
		return
	}

	balance, err := schoolcardService.GetCardToday(user)
	if err == errors.ERR_Session_Expired {
		user, err = controller.LoginHandle(context, schoolcardService.GetUser)
		balance, err = schoolcardService.GetCardToday(user)
	}

	if err != nil {
		controller.ErrorHandle(context, err)
		return
	}
	utils.ContextDataResponseJson(context, utils.SuccessResponseJson(balance))
}

// @Summary 校园卡历史查询
// @Description 校园卡历史查询
// @Tags 校园卡
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param year body string true "年份"
// @Param month body string true "月份"
// @Success 200 json  {"code":200,"data":[{...}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/card/history [post]
func CardHistory(context *gin.Context) {
	isValid := utils.CheckPostFormEmpty(
		context,
		[]string{"username", "password", "year", "month"},
	)

	if !isValid {
		utils.ContextDataResponseJson(context, utils.FailResponseJson(errors.InvalidArgs, nil))
		return
	}

	user, err := controller.LoginHandle(context, schoolcardService.GetUser)
	if err != nil {
		return
	}

	history, err := schoolcardService.GetCardHistory(user, context.PostForm("year"), context.PostForm("month"))
	if err == errors.ERR_Session_Expired {
		user, err = controller.LoginHandle(context, schoolcardService.GetUser)
		history, err = schoolcardService.GetCardHistory(user, context.PostForm("year"), context.PostForm("month"))
	}
	if err != nil {
		controller.ErrorHandle(context, err)
		return
	}
	utils.ContextDataResponseJson(context, utils.SuccessResponseJson(history))
}
