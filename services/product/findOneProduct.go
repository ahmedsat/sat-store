package product

import (
	"strings"

	"github.com/ahmedsat/sat-store/services"
)

func (p *Product) FindOne(data map[string]string) (services.Entity, error) {

	searchConditions := services.SearchMapParser(data)

	var dbImage string

	row := services.DB.QueryRow("SELECT * FROM products " + searchConditions)

	if err := row.Scan(&p.Id, &p.Name, &p.Categories, &p.Price, &p.Discount, &p.Description, &dbImage); err != nil {
		return nil, err
	}
	p.Images = strings.Split(dbImage, ",")
	println(p.Images[0])
	return p, nil
}

func (p *Product) Delete(data map[string]string) (bool, error) {

	searchConditions := services.SearchMapParser(data)

	// var dbImage string

	_, err := services.DB.Exec("DELETE FROM products " + searchConditions + " LIMIT 1")
	if err != nil {
		return false, err
	}

	return true, nil
}
