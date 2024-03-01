package helpers

import (
	"restAPI/database/models"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
