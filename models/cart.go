package models

import "gorm.io/gorm"

type CartItems struct {
	gorm.Model

	Product Product `json:"Product"`
	User    User    `json:"user"`

	Count  uint   `json:"count"`
	Status string `sql:"type:ENUM('waiting', 'packaging','shipping',received)" json:"status"`
}
