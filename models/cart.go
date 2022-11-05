package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model

	Product Product `json:"Product" gorm:"foreignKey:ID;"`
	User    User    `json:"user" gorm:"foreignKey:ID;"`
	Count   uint    `json:"count"`

	Status string `sql:"type:ENUM('onCart','ordered')" json:"status"`
}
