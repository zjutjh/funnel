package router

import (
	"funnel/app/controller/canteenController"
	"funnel/app/controller/libraryController"
	"funnel/app/controller/schoolCardController"
	"funnel/app/controller/zfController"
	"funnel/app/midware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	student := r.Group("/student")
	{
		zf := student.Group("/zf", midware.CheckUsernamePassword)
		{
			term := zf.Group("", midware.CheckTermInfoForm)
			{
				term.POST("/score/info", zfController.GetScoreDetail)
				term.POST("/score", zfController.GetScore)
				term.POST("/midtermscore", zfController.GetMidTermScore)
				term.POST("/table", zfController.GetClassTable)
				term.POST("/exam", zfController.GetExamInfo)
			}
			zf.POST("/room", zfController.GetRoomInfo)
			zf.POST("/program", zfController.GetProgInfo)
		}
		library := student.Group("/library", midware.CheckUsernamePassword)
		{
			library.POST("/borrow/history", libraryController.LibraryBorrowHistory)
			library.POST("/borrow/current", libraryController.LibraryCurrentBorrow)
			library.POST("/borrow/reborrow", libraryController.LibraryReBorrow)
		}
		card := student.Group("/card", midware.CheckUsernamePassword)
		{
			card.POST("/balance", schoolCardController.CardBalance)
			card.POST("/today", schoolCardController.CardToday)
			card.POST("/history", schoolCardController.CardHistory)
		}
	}
	canteen := r.Group("/canteen")
	{
		canteen.GET("/flow", canteenController.Flow) // 关于餐厅客流量的路由
	}

	return r
}
