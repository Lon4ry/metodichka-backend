package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"metodichka-backend/models"
	"net/http"
)

func GetCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.Category
		if err := db.Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}
