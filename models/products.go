package models

type Product struct {
	CustomModel

	Images   []Image  `json:"images" gorm:"foreignKey:ID;"`
	Category Category `json:"category" gorm:"foreignKey:ID;"`

	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Discount    float64 `json:"discount"`
}
