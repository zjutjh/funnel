package router

import (
	"funnel/app/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	student := r.Group("/student")
	{
		zf := student.Group("/zf")
		{
			zf.POST("/score/info", controller.GetScoreDetail)
			zf.POST("/score", controller.GetScore)
			zf.POST("/table", controller.GetClassTable)
			zf.POST("/exam", controller.GetExamInfo)
			zf.POST("/room", controller.GetRoomInfo)
			zf.POST("/program", controller.GetProgInfo)
		}
		library := student.Group("/libraryService")
		{
			library.POST("/borrow/history/:page", controller.LibraryBorrowHistory)
			library.POST("/borrow/current", controller.LibraryCurrentBorrow)
		}
		card := student.Group("/card")
		{
			card.POST("/balance", controller.CardBalance)
			card.POST("/today", controller.CardToday)
			card.POST("/history", controller.CardHistory)
		}
	}

	return r
}
