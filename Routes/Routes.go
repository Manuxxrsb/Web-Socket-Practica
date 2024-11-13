package routes

import (
	database "backend/Database"
	handlers "backend/Handlers"
	models "backend/Models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func RunServer(router *gin.Engine) {

	db, err := database.OpenGormDB()

	router.GET("/ping", handlers.Ping())
	if err != nil {
		log.Fatalf("Error al conectarse a la BD: %v", err)
	}

	db.AutoMigrate(&models.Mensaje{})
	db.AutoMigrate(&models.Canal{})

	fmt.Print("Conexión exitosa a la BD")

	//Rutas para el Web Socket
	router.GET("/Web_Socket", handlers.Ejecuta_Socket())

	//Rutas para el CRUD de Mensajes
	router.POST("/Mensaje", handlers.Create_Mensaje(db))
	router.GET("/Mensaje/:id", handlers.Get_Mensaje(db))
	router.GET("/Mensajes", handlers.Getall_Mensaje(db))
	router.PUT("/Mensaje/:id", handlers.Put_Mensaje(db))
	router.DELETE("/Mensaje/:id", handlers.Delete_Mensaje(db))

	//Rutas para el CRUD de Canal

	router.POST("/Canal", handlers.Create_Canal(db))
	router.GET("/Canal/:id", handlers.Get_Canal(db))
	router.GET("/Canales", handlers.Getall_Canal(db))
	router.PUT("/Canal/:id", handlers.Put_Canal(db))
	router.DELETE("/Canal/:id", handlers.Delete_Canal(db))

	fmt.Println("El servidor HTTP está ejecutándose en :8080")
}
