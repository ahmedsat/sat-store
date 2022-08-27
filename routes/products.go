package routes

import (
	"database/sql"
	"net/http"

	"github.com/ahmedsat/sat-store/resources"
	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func GetAllProducts(c *gin.Context) {

	resources.SetDB(DB)

	p := resources.Entity.New(&resources.Product{})

	products, err := p.Find(map[string]string{})
	if err != nil {
		panic(err)
	}

	result := struct {
		Size int
		Data []resources.Entity
	}{len(products), products}
	c.IndentedJSON(http.StatusOK, result)

}

func AddProduct(c *gin.Context) {

	resources.SetDB(DB)

	var product *resources.Product
	if err := c.BindJSON(&product); err != nil {
		panic(err)
	}

	if _, err := product.Create(); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, product)
}
