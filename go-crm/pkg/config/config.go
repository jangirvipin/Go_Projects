package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DB *gorm.DB
)

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dsn := os.Getenv("DB_DSN")
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
