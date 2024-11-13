package handlers

import (
	models "backend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Put_Canal(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		id := Request.Param("id")
		var Canal models.Canal
		if err := db.First(&Canal, id).Error; err != nil {
			Request.JSON(http.StatusNotFound, gin.H{"error": "Canal no encontrado " + err.Error()})
			return
		}

		var input models.Canal
		if err := Request.ShouldBindJSON(&input); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&Canal).Updates(input).Error; err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar El Canal en la BD " + err.Error()})
			return
		}

		Request.JSON(http.StatusOK, Canal)
	}
}
