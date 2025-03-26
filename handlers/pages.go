package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"metodichka-backend/models"
	"net/http"
)

func GetPage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page models.Page
		if err := db.Find(&page, "category_id = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, page)
	}
}
