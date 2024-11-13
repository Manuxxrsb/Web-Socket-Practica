package handlers

import (
	models "backend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Getall_Mensaje(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		var Mensaje []models.Mensaje
		err := db.Find(&Mensaje).Error //todos los registros que coinciden con el modelo

		if err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"error": "No existen Mensajes " + err.Error()})
			return
		}
		Request.JSON(http.StatusOK, Mensaje)
	}
}

func Get_Mensaje(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		id := Request.Param("id")
		var Mensaje models.Mensaje
		if err := db.First(&Mensaje, id).Error; err != nil { //Primero encuentra el primer registro ordenado por clave principal
			Request.JSON(http.StatusNotFound, gin.H{"error": "Mensaje no encontrado " + err.Error()})
			return
		}
		Request.JSON(http.StatusOK, Mensaje)
	}
}
