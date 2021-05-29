package database

import (
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/donreno/gofiber-test-api/config"
)

func Connect() (db *gorm.DB, err error) {
	p := config.Get("DB_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		return nil, fmt.Errorf("error converting port value: %s to int", p)
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		config.Get("DB_HOST"),
		port,
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"))

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %w", err)
	}

	return db, nil
}
