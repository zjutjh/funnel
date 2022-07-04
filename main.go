package main

import (
	"funnel/config"
	"funnel/router"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	config.SetupConfigs(r)
	router.SetupRouter(r)
	_ = r.Run(":" + os.Getenv("ROUTER_POST"))
}
