package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	db, err := gorm.Open(sqlite.Open("database/gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db

}
