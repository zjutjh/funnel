package libraryController

import (
	"funnel/app/controller"
	"funnel/app/service/libraryService"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

type currentBorrowData struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Page     int    `form:"page" json:"page"  default:"1"`
}

// LibraryCurrentBorrow 图书馆当前借书记录
func LibraryCurrentBorrow(c *gin.Context) {
	var data currentBorrowData
	if err := c.ShouldBind(&data); err != nil {
		controller.ErrorHandle(c, err)
		return
	}
	currentBorrow, err := libraryService.GetCurrentBorrow(data.Username, data.Password, data.Page)
	if err != nil {
		controller.ErrorHandle(c, err)
		return
	}
	utils.ContextDataResponseJson(c, utils.SuccessResponseJson(currentBorrow))
}
