package main

import (
	"encoding/json"
	funnel "funnel/app"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"time"
)

func RedisInit() *redis.Client {
	RedisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
	return RedisClient
}

func main() {
	ts := funnel.TeachingAdministrationSystem{RootUrl: "http://172.16.19.163/jwglxt/"}
	r := gin.Default()

	redisX := RedisInit()

	r.Group("/v2")
	{
		r.Group("/student")
		{
			r.POST("/login", func(c *gin.Context) {

				name := c.PostForm("username")
				passwd := c.PostForm("password")
				stu := funnel.Student{Sid: name, Password: passwd}

				for {
					err := ts.Login(&stu)
					if err == nil {
						break
					}
				}
				ser, _ := json.Marshal(stu)
				redisX.Set(stu.Sid, string(ser), time.Duration(stu.Session.MaxAge))
				redisX.Save()

				c.Data(200, "application/json", []byte("OK"))
			})

			r.GET("/score-info/:username/:year/:term", func(c *gin.Context) {
				stuSer, _ := redisX.Get(c.Param("username")).Result()
				stu := &funnel.Student{}
				_ = json.Unmarshal([]byte(stuSer), stu)
				c.Data(200, "application/json", []byte(ts.GetScoreDetail(stu, c.Param("year"), c.Param("term"))))
			})
			r.GET("/score/:username/:year/:term", func(c *gin.Context) {
				stuSer, _ := redisX.Get(c.Param("username")).Result()
				stu := &funnel.Student{}
				_ = json.Unmarshal([]byte(stuSer), stu)
				c.Data(200, "application/json", []byte(ts.GetScore(stu, c.Param("year"), c.Param("term"))))
			})

			r.GET("/class/:username/:year/:term", func(c *gin.Context) {
				stuSer, _ := redisX.Get(c.Param("username")).Result()
				stu := &funnel.Student{}
				_ = json.Unmarshal([]byte(stuSer), stu)
				c.Data(200, "application/json", []byte(ts.GetClassTable(stu, c.Param("year"), c.Param("term"))))
			})
			r.GET("/exam/info/:username/:year/:term", func(c *gin.Context) {
				stuSer, _ := redisX.Get(c.Param("username")).Result()
				stu := &funnel.Student{}
				_ = json.Unmarshal([]byte(stuSer), stu)
				c.Data(200, "application/json", []byte(ts.GetExamInfo(stu, c.Param("year"), c.Param("term"))))
			})
		}

	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
