package migrations

import (
	"gin-gorm/models"

	"gorm.io/gorm"
)

func InitializeMigration(db *gorm.DB) error {
	return db.AutoMigrate(models.User{})
}
