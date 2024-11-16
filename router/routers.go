package router

import (
	"funnel/app/controller/libraryController"
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
		library := student.Group("/library")
		{
			library.POST("/borrow/history", libraryController.LibraryBorrowHistory)
			library.POST("/borrow/current", libraryController.LibraryCurrentBorrow)
			library.POST("/borrow/reborrow", libraryController.LibraryReBorrow)
		}
	}

	return r
}
