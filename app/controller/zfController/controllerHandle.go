package zfController

import (
	"encoding/json"
	"funnel/app/controller"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service/zfService"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

func ZFTermInfoHandle(context *gin.Context, cb func(*model.User, string, string) (string, error)) (string, error) {
	user, err := controller.LoginHandle(context, zfService.GetUser)
	if err != nil {
		return "", err
	}

	result, err := cb(user, context.PostForm("year"), context.PostForm("term"))

	if err == errors.ERR_Session_Expired {
		user, err = controller.LoginHandle(context, zfService.GetUser)
		result, err = cb(user, context.PostForm("year"), context.PostForm("term"))
	}

	if err != nil {
		controller.ErrorHandle(context, err)
		return "", err
	}

	var f interface{}
	err = json.Unmarshal([]byte(result), &f)
	if err != nil {
		controller.ErrorHandle(context, err)
		return "", err
	}
	utils.ContextDataResponseJson(context, utils.SuccessResponseJson(f))
	return result, err
}
