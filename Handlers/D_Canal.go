package handlers

import (
	models "backend/Models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete_Canal(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id " + err.Error()})
			return
		}

		var Canal models.Canal
		if err := db.First(&Canal, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No se pudo encontrar la persona para eliminar " + err.Error()})
			return
		}

		if err := db.Delete(&Canal).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la Canal " + err.Error()})
			return
		}
		//delete no elimina la tabla pero para gorm llena deletedat que lo ignorarara para los find o first

		c.JSON(http.StatusOK, gin.H{"message": "Canal eliminado con éxito "})
	}
}
