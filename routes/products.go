package routes

import (
	"database/sql"
	"net/http"

	"github.com/ahmedsat/sat-store/resources"
	"github.com/gin-gonic/gin"
)

// type product struct {
// 	Id          int `json:"id"`
// 	Name        string
// 	Categories  string
// 	Price       string
// 	Discount    string
// 	Description string
// 	Images      []string
// }

var DB *sql.DB

func GetAllProducts(c *gin.Context) {

	resources.SetDB(DB)

	// var products []resources.Product

	p := resources.Entity.New(&resources.Product{})

	p.Find(map[string]any{"Name": "Sat", "id": "id"})

	// rows, err := DB.Query("SELECT * FROM products")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	var p product
	// 	var dbImage string
	// 	if err = rows.Scan(&p.Id, &p.Name, &p.Categories, &p.Price, &p.Discount, &p.Description, &dbImage); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	p.Images = strings.Split(dbImage, ",")
	// 	products = append(products, p)
	// }

	c.IndentedJSON(http.StatusOK, "result")

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
