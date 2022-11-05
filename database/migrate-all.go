package database

import "github.com/ahmedsat/sat-store/models"

func MigrateAll() {

	DB.AutoMigrate(
		&models.Image{},
		&models.Category{},
		&models.User{},
		&models.Product{},
		&models.CartItem{},
		&models.Order{},
	)

}
