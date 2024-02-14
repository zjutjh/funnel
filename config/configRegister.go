package config

import (
	"funnel/config/logs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func SetupConfigs(r *gin.Engine) {
	log.Print("Setup Configs....")
	RedisInit()
	logs.LogInit()
	log.Print("Setup Configs Finish....")
}
