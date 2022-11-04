package models

import (
	"gorm.io/gorm"
)

type CustomModel struct {
	gorm.Model
	Name string `json:"name"`
}
