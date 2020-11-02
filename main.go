package main

import (
	"funnel/app/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("sessions", store))

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

			}
		}
		r.Group("/teacher")
		{

		}

	}

	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
