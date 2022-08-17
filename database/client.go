package database

import (
	"go-api-jwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
	}
	log.Println("Connect to Database!")
}

func Migrate() {
	err := Instance.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	log.Println("Database Migration Completed!")
}
