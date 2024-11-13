package handlers

import (
	models "backend/Models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete_Mensaje(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id " + err.Error()})
			return
		}

		var Mensaje models.Mensaje
		if err := db.First(&Mensaje, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No se pudo encontrar la persona para eliminar " + err.Error()})
			return
		}

		if err := db.Delete(&Mensaje).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la Mensaje " + err.Error()})
			return
		}
		//delete no elimina la tabla pero para gorm llena deletedat que lo ignorarara para los find o first

		c.JSON(http.StatusOK, gin.H{"message": "Mensaje eliminado con éxito "})
	}
}
