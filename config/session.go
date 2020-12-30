package config

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"os"
)
import "github.com/gin-gonic/gin"

func setupSession(r *gin.Engine)  {
	store, _ := redis.NewStore(10, "tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASSWORD"), []byte("secret"))
	r.Use(sessions.Sessions("sessions", store))
}
