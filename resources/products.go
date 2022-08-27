package resources

import (
	"fmt"
	"strings"
)

type Product struct {
	Id          int `json:"id"`
	Name        string
	Categories  string
	Price       float64
	Discount    float64
	Description string
	Images      []string
}

func (p *Product) Create() (Entity, error) {

	add, err := db.Prepare(`--sql
	insert into products (name,categories,price,discount,description,images)
	values (?,?,?,?,?,?)
	`)
	if err != nil {
		return nil, err
	}
	defer add.Close()

	dbImages := strings.Join(p.Images, ",")
	result, err := add.Exec(p.Name, p.Categories, p.Price, p.Discount, p.Description, dbImages)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	p.Id = int(id)

	return p, nil

}

func (p *Product) Find(data map[string]any) ([]Entity, error) {

	search := make(map[string]any)

	if sId, ok := data["Id"]; ok {
		search["Id"] = sId
	}

	if sName, ok := data["Name"]; ok {
		search["Name"] = sName
	}

	if sCategories, ok := data["Categories"]; ok {
		search["Categories"] = sCategories
	}

	if sPrice, ok := data["Price"]; ok {
		search["Price"] = sPrice
	}

	if sDiscount, ok := data["Discount"]; ok {
		search["Discount"] = sDiscount
	}

	if sDescription, ok := data["Description"]; ok {
		search["Description"] = sDescription
	}

	if sImages, ok := data["Images"]; ok {
		search["Images"] = sImages
	}

	// rows, err := DB.Query("SELECT * FROM products")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return nil, nil
}

func (p *Product) New() Entity {
	return &Product{}
}
