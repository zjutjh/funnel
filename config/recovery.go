package config

import (
	"log"
)

func setupRecover() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
}
