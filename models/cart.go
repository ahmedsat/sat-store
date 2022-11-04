package models

import "gorm.io/gorm"

type CartItems struct {
	gorm.Model

	Product Product `json:"Product" gorm:"foreignKey:ID;"`
	User    User    `json:"user" gorm:"foreignKey:ID;"`

	Count  uint   `json:"count"`
	Status string `sql:"type:ENUM('waiting', 'packaging','shipping',received)" json:"status"`
}
