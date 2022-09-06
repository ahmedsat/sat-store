package product

import "github.com/ahmedsat/sat-store/services"

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Categories  string  `json:"categories"`
	Price       float64 `json:"price,string"`
	Discount    float64 `json:"discount,string,omitempty"`
	Description string  `json:"Description,omitempty"`
	Images      []string
}

func (p *Product) New() services.Entity {
	return p
}