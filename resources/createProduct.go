package resources

import "strings"

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
