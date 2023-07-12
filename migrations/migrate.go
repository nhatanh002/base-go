package migrations

import (
	"base-go/common/logger"
	"base-go/domain/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	logger.Info("Running migrations for Cat model...")
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Cat{})
}
