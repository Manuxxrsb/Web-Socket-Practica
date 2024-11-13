package handlers

import (
	models "backend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Getall_Canal(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		var Canal []models.Canal
		err := db.Find(&Canal).Error //todos los registros que coinciden con el modelo

		if err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"error": "No existen Canals " + err.Error()})
			return
		}
		Request.JSON(http.StatusOK, Canal)
	}
}

func Get_Canal(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		id := Request.Param("id")
		var Canal models.Canal
		if err := db.First(&Canal, id).Error; err != nil { //Primero encuentra el primer registro ordenado por clave principal
			Request.JSON(http.StatusNotFound, gin.H{"error": "Canal no encontrado " + err.Error()})
			return
		}

		//agregamos la busqueda para los objetos de la Persona
		db.Model(&Canal).Association("Mensajes").Find(&Canal.Mensajes)

		Request.JSON(http.StatusOK, Canal)
	}
}
