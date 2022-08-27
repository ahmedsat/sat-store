package resources

import (
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

func (p *Product) Find(data map[string]string) ([]Entity, error) {

	var products []Entity

	searchMap := make(map[string]string)
	searchConditions := ""

	if sId, ok := data["Id"]; ok {
		searchMap["id"] = sId
	}

	if sName, ok := data["Name"]; ok {
		searchMap["name"] = sName
	}

	if sCategories, ok := data["Categories"]; ok {
		searchMap["categories"] = sCategories
	}

	if sPrice, ok := data["Price"]; ok {
		searchMap["price"] = sPrice
	}

	if sDiscount, ok := data["Discount"]; ok {
		searchMap["discount"] = sDiscount
	}

	if sDescription, ok := data["Description"]; ok {
		searchMap["description"] = sDescription
	}

	if sImages, ok := data["Images"]; ok {
		searchMap["images"] = sImages
	}

	if len(searchMap) > 0 {
		searchConditions += "WHERE "
		for k, v := range searchMap {
			searchConditions += k + " = '" + v + "' AND "
		}
	}
	searchConditions = strings.TrimSuffix(searchConditions, "AND ")
	println(searchConditions)

	rows, err := db.Query("SELECT * FROM products " + searchConditions)
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

func (p *Product) New() Entity {
	return &Product{}
}
