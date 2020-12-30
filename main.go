package main

import (
	"funnel/config"
	"funnel/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.SetupConfigs(r)
	router.SetupRouter(r)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
