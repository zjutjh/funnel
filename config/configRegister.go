package config

import (
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
	setupSession(r)
	log.Print("Setup Configs Finish....")
}
