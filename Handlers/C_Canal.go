package handlers

import (
	models "backend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create_Canal(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		var Canal models.Canal
		if err := Request.ShouldBindJSON(&Canal); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&Canal).Error; err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		Request.JSON(http.StatusOK, Canal)
	}

}
