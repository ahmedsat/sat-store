package models

import "gorm.io/gorm"

type CartItems struct {
	gorm.Model

	Product Product `json:"Product"`
	User    User    `json:"user"`

	Count uint `json:"count"`
}
