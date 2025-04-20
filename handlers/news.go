package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"metodichka-backend/models"
	"net/http"
	"strings"
	"time"
)

type News struct {
	Title    string `json:"title"  validate:"required"`
	Author   string `json:"author"  validate:"required"`
	Markdown string `json:"markdown"  validate:"required "`
	Img      string `json:"img"  validate:"required "`
}

func GetNews(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var news []models.News
		if err := db.Find(&news).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, news)
	}
}
func SetNews(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req News

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()

		if err := validate.Struct(req); err != nil {
			time.Sleep(1 * time.Second)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		news := models.News{
			Author:   strings.Trim(req.Author, " "),
			Title:    strings.Trim(req.Title, " "),
			Markdown: strings.Trim(req.Markdown, " "),
			Img:      strings.Trim(req.Img, " "),
		}

		if err := db.Create(&news).Error; err != nil {
			time.Sleep(1 * time.Second)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": news})
	}
}
