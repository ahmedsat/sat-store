package product

import "github.com/ahmedsat/sat-store/services"

func (p *Product) Delete(data map[string]string) (bool, error) {

	searchConditions := services.SearchMapParser(data)

	// var dbImage string

	_, err := services.DB.Exec("DELETE FROM products " + searchConditions + " LIMIT 1")
	if err != nil {
		return false, err
	}

	return true, nil
}
