package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("DATABASE_URL")

	d, err2 := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err2 != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
