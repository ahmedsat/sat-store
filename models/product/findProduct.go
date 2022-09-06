package product

import (
	"strings"

	"github.com/ahmedsat/sat-store/services"
)

func (p *Product) Find(data map[string]string) ([]services.Entity, error) {

	var products []services.Entity

	searchConditions := services.SearchMapParser(data)

	rows, err := services.DB.Query("SELECT * FROM products " + searchConditions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		var dbImage string
		if err = rows.Scan(&p.Id, &p.Name, &p.Categories, &p.Price, &p.Discount, &p.Description, &dbImage); err != nil {
			return nil, err
		}
		p.Images = strings.Split(dbImage, ",")
		products = append(products, &p)
	}

	return products, nil
}
