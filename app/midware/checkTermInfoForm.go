package midware

import (
	"funnel/app/errors"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

func CheckTermInfoForm(context *gin.Context) {
	isValid := utils.CheckPostFormEmpty(
		context,
		[]string{"username", "password", "year", "term"},
	)
	if !isValid {
		utils.ContextDataResponseJson(context, utils.FailResponseJson(errors.InvalidArgs, nil))
		context.Abort()
	}
}
