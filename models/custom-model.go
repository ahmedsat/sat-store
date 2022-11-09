package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomModel struct {
	ID        uint           `gorm:"primarykey" form:"id"`
	CreatedAt time.Time      `form:"created_at"`
	UpdatedAt time.Time      `form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" form:"deleted_at"`
	Name      string         `json:"name" gorm:"default:no name" form:"name"`
}
