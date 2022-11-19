package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"

	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_DATABASE")
	DbDriver := os.Getenv("DB_DRIVER")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbName)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to the database", DbDriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database", DbDriver)
	}

	DB.AutoMigrate(&User{})
}
