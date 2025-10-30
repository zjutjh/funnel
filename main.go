package main

import (
	"funnel/app/apis"
	"funnel/app/service/zfService/loginTokenManager"
	"funnel/config/config"
	"funnel/router"

	"github.com/gin-gonic/gin"
)

func main() {
	loginTokenManager.Init(apis.ZF_HOST)
	r := gin.Default()
	router.SetupRouter(r)
	_ = r.Run(":" + config.Config.GetString("port"))
}
