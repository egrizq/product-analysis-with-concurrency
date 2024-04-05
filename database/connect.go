package database

import (
	"fmt"
	"log"
	"os"
	"product-store-management/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializationDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, database, password)

	connect, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	return connect
}

func Connection() {
	DB = InitializationDB()
	if err := DB.AutoMigrate(&model.Product{}); err != nil {
		log.Fatalf("Error migrating database schema: %v", err)
	}
}
