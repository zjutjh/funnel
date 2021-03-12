package canteenController

import (
	"funnel/app/errors"
	"funnel/app/service/canteenService"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

func PeopleFlow(content *gin.Context) {
	data, isSuccessed := canteenService.FetchFlow()
	if isSuccessed {
		utils.ContextDataResponseJson(
			content,
			utils.SuccessResponseJson(data))
	} else {
		utils.ContextDataResponseJson(
			content,
			utils.FailResponseJson(errors.RequestFailed, data))
	}
}
