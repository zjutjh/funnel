package main

import (
	"funnel/config/config"
	"funnel/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	_ = r.Run(":" + config.Config.GetString("port"))
}
