package handlers

import (
	models "backend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Put_Mensaje(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		id := Request.Param("id")
		var Mensaje models.Mensaje
		if err := db.First(&Mensaje, id).Error; err != nil {
			Request.JSON(http.StatusNotFound, gin.H{"error": "Mensaje no encontrado " + err.Error()})
			return
		}

		var input models.Mensaje
		if err := Request.ShouldBindJSON(&input); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&Mensaje).Updates(input).Error; err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar El Mensaje en la BD " + err.Error()})
			return
		}

		Request.JSON(http.StatusOK, Mensaje)
	}
}
