package models

type Product struct {
	CustomModel

	Images   []Image  `json:"images"`
	Category Category `json:"category"`

	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Discount    float64 `json:"discount"`
}
