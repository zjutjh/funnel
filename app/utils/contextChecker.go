package utils

import (
	"github.com/gin-gonic/gin"
)

func CheckPostFormEmpty(context *gin.Context, items []string) bool {
	for i := range items {
		_, has := context.GetPostForm(items[i])
		if has == false {
			return false
		}
	}
	return true
}
