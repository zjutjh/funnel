package libraryController

import (
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

func LibraryReBorrow(c *gin.Context) {
	utils.ContextDataResponseJson(c, utils.ResponseJsonMessage{
		Code:    500,
		Message: "功能仍未上架,敬请期待",
		Data:    interface{}(nil),
	})
}
