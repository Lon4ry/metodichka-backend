package db_config

import (
	"gorm.io/gorm"
	"metodichka-backend/models"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Category{}, &models.Page{}, &models.News{})
	if err != nil {
		return err
	}
	return nil
}
