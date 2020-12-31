package router

import (
	"funnel/app/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	v2 := r.Group("/v2")
	{
		student := v2.Group("/student")
		{
			zf := student.Group("/zf")
			{
				zf.POST("/login", controller.ZFLogin)
				zf.POST("/score/info", controller.GetScoreDetail)
				zf.POST("/score", controller.GetScore)
				zf.POST("/table", controller.GetClassTable)
				zf.POST("/exam", controller.GetExamInfo)
			}
			library := student.Group("/library")
			{
				library.POST("/login", controller.LibraryLogin)
				library.GET("/borrow/history/:page", controller.LibraryBorrowHistory)
				library.GET("/borrow/current/:page", controller.LibraryCurrentBorrow)
			}
			card := student.Group("/card")
			{
				card.POST("/login", controller.CardLogin)
				card.Any("/balance", controller.CardBalance)
				card.Any("/today", controller.CardToday)
			}
		}
	}

	return r
}
