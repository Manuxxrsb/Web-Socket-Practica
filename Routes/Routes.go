package routes

import (
	handlers "backend/Handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer(router *gin.Engine) {

	router.GET("/ping", handlers.Ping())
	router.GET("/Web_Socket", handlers.Ejecuta_Socket())

	fmt.Println("El servidor HTTP está ejecutándose en :8080")
}
