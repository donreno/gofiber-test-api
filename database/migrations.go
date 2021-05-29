package database

import (
	"fmt"

	"github.com/donreno/gofiber-test-api/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Product{}); err != nil {
		return fmt.Errorf("error automigrating model %w", err)
	}

	return nil
}
