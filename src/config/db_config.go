package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "crud/src/model"
) 

var DB *gorm.DB

func InitDB() {
	// Load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Failed to load .env file")
	}

	// Get DSN from environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL not set in environment")
	}

	// Open PostgreSQL connection (for Neon)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// AutoMigrate models
	err = db.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Product{},
		&models.PaymentHistory{},
	)
	if err != nil {
		panic(fmt.Sprintf("AutoMigrate failed: %v", err))
	}

	DB = db
}
