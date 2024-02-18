package main

import (
	"funnel/router"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	_ = r.Run(":" + os.Getenv("ROUTER_POST"))
}
