package router

import (
	"funnel/app/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	r.Group("/v2")
	{
		r.Group("/student")
		{
			r.Group("/zf")
			{
				r.POST("/login", controller.ZFLogin)
				r.GET("/score-info/:year/:term", controller.GetScoreDetail)
				r.GET("/score/:year/:term", controller.GetScore)
				r.GET("/class-table/:year/:term", controller.GetClassTable)
				r.GET("/exam/:username/:year/:term", controller.GetExamInfo)
			}
			r.Group("/library")
			{
				r.POST("/library/login", controller.LibraryLogin)
				r.GET("/borrow/history/:page", controller.LibraryBorrowHistory)
				r.GET("/borrow/current/:page", controller.LibraryCurrentBorrow)
			}
		}
		r.Group("/teacher")
		{

		}

	}

	return r
}