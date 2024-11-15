package libraryController

import (
	"funnel/app/controller"
	"funnel/app/service/libraryService"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

type borrowHistoryData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Page     int    `json:"page" default:"1"`
}

// LibraryBorrowHistory 图书馆历史借书记录
func LibraryBorrowHistory(c *gin.Context) {
	var data borrowHistoryData
	if err := c.ShouldBind(&data); err != nil {
		controller.ErrorHandle(c, err)
		return
	}
	currentBorrow, err := libraryService.GetBorrowHistory(data.Username, data.Password, data.Page)
	if err != nil {
		controller.ErrorHandle(c, err)
		return
	}
	utils.ContextDataResponseJson(c, utils.SuccessResponseJson(currentBorrow))
}
