package canteenController

import (
	"fmt"
	"funnel/app/errors"
	"funnel/app/service/canteenService"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

func Flow(content *gin.Context) {
	data, err := canteenService.FetchFlow()
	if err == nil {
		utils.ContextDataResponseJson(
			content,
			utils.SuccessResponseJson(data))
	} else {
		utils.ContextDataResponseJson(
			content,
			utils.FailResponseJson(errors.RequestFailed, fmt.Sprintf("%v: %v", data, err)))
	}
}
