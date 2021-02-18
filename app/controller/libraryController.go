package controller

import (
	"funnel/app/errors"
	"funnel/app/service/libraryService"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 图书馆历史借书记录
// @Description 图书馆借书记录（暂时只支持10本）
// @Tags 图书馆
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":[{...}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/libraryService/history/0 [post]
func LibraryBorrowHistory(context *gin.Context) {
	isValid := utils.CheckPostFormEmpty(
		context,
		[]string{"username", "password"},
	)

	if !isValid {
		utils.ContextDataResponseJson(context, utils.FailResponseJson(errors.InvalidArgs, nil))
		return
	}

	user, err := libraryService.GetUser(context.PostForm("username"), context.PostForm("password"))
	if err == errors.ERR_WRONG_PASSWORD {
		utils.ContextDataResponseJson(context, utils.FailResponseJson(errors.WrongPassword, nil))
		return
	}
	if err != nil {
		utils.ContextDataResponseJson(context, utils.FailResponseJson(errors.UnKnown, nil))
		return
	}
	books := libraryService.GetBorrowHistory(user)
	utils.ContextDataResponseJson(context, utils.SuccessResponseJson(books))
}

// @Summary 图书馆当前借书记录
// @Description 图书馆当前借书记录
// @Tags 图书馆
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":[{...}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/libraryService/current [post]
func LibraryCurrentBorrow(context *gin.Context) {
	isValid := utils.CheckPostFormEmpty(
		context,
		[]string{"username", "password"},
	)

	if !isValid {
		utils.ContextDataResponseJson(context, utils.FailResponseJson(errors.InvalidArgs, nil))
		return
	}

	user, err := libraryService.GetUser(context.PostForm("username"), context.PostForm("password"))
	if err == errors.ERR_WRONG_PASSWORD {
		utils.ContextDataResponseJson(context, utils.FailResponseJson(errors.WrongPassword, nil))
		return
	}
	if err != nil {
		utils.ContextDataResponseJson(context, utils.FailResponseJson(errors.UnKnown, nil))
		return
	}

	books := libraryService.GetCurrentBorrow(user)
	utils.ContextDataResponseJson(context, utils.SuccessResponseJson(books))
}
