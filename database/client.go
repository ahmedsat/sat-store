package database

import (
	"log"

	"github.com/ahmedsat/sat-store/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(database string) {
	Instance, dbError = gorm.Open(sqlite.Open(database), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	// Instance.AutoMigrate(&product.Product{})
	log.Println("Database Migration Completed!")
}
