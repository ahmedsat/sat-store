package resources

import "strings"

func (p *Product) FindOne(data map[string]string) (Entity, error) {

	searchConditions := searchMapParser(data)

	var dbImage string

	row := db.QueryRow("SELECT * FROM products " + searchConditions)
	if err := row.Scan(&p.Id, &p.Name, &p.Categories, &p.Price, &p.Discount, &p.Description, &dbImage); err != nil {
		return nil, err
	}
	p.Images = strings.Split(dbImage, ",")
	println(p.Images[0])
	return p, nil
}
