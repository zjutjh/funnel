package helps

import (
	"github.com/gin-gonic/gin"
)
import "funnel/app/errors"

type ResponseJsonMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponseJson(d interface{}) ResponseJsonMessage {
	return ResponseJsonMessage{Code: 200, Message: "OK", Data: d}
}

func FailResponseJson(error errors.ResponseError, d interface{}) ResponseJsonMessage {
	return ResponseJsonMessage{Code: error.Code, Message: error.Message, Data: d}
}

func ContextDataResponseJson(context *gin.Context, response ResponseJsonMessage) {
	context.JSON(200, response)
}
