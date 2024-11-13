package handlers

import "github.com/gin-gonic/gin"

func Ping() gin.HandlerFunc {
	return func(Request *gin.Context) {
		Request.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
