package migrations

import (
	"restAPI/database/schema/models"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
