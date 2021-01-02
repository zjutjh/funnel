package controller

import (
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/service"
	"github.com/gin-gonic/gin"
)

var librarySystem = service.LibrarySystem{}

// @Summary 图书馆历史借书记录
// @Description 图书馆借书记录（暂时只支持10本）
// @Tags 图书馆
// @Produce  json
// @Success 200 json  {"code":200,"data":[{...}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/library/history/0 [get]
func LibraryBorrowHistory(context *gin.Context) {
	isValid := helps.CheckPostFormEmpty(
		context,
		[]string{"username", "password"},
	)

	if !isValid {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.InvalidArgs, nil))
		return
	}

	user, err := librarySystem.GetUser(context.PostForm("username"), context.PostForm("password"))
	if err == errors.ERR_WRONG_PASSWORD {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.WrongPassword, nil))
		return
	}
	if err != nil {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
		return
	}
	books := librarySystem.GetBorrowHistory(user)
	helps.ContextDataResponseJson(context, helps.SuccessResponseJson(books))
}

// @Summary 图书馆当前借书记录
// @Description 图书馆当前借书记录
// @Tags 图书馆
// @Produce  json
// @Success 200 json  {"code":200,"data":[{...}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/library/current/0 [get]
func LibraryCurrentBorrow(context *gin.Context) {
	isValid := helps.CheckPostFormEmpty(
		context,
		[]string{"username", "password"},
	)

	if !isValid {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.InvalidArgs, nil))
		return
	}

	user, err := librarySystem.GetUser(context.PostForm("username"), context.PostForm("password"))
	if err == errors.ERR_WRONG_PASSWORD {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.WrongPassword, nil))
		return
	}
	if err != nil {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
		return
	}

	books := librarySystem.GetCurrentBorrow(user)
	helps.ContextDataResponseJson(context, helps.SuccessResponseJson(books))
}
