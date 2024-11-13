package handlers

import (
	models "backend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create_Mensaje(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		var Mensaje models.Mensaje
		if err := Request.ShouldBindJSON(&Mensaje); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&Mensaje).Error; err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		Request.JSON(http.StatusOK, Mensaje)
	}

}
