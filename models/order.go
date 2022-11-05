package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	Phone   string `json:"Phone"`
	Address string `json:"address"`

	CartItems []CartItem `json:"cartItems" gorm:"foreignKey:ID;"`

	Status string `sql:"type:ENUM('waiting', 'packaging','shipping','done')" json:"status"`
}
