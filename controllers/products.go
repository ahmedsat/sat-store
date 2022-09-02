package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ahmedsat/sat-store/services"
	"github.com/ahmedsat/sat-store/services/product"
	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func GetAllProducts(c *gin.Context) {

	services.SetDB(DB)

	p := services.Entity.New(&product.Product{})

	products, err := p.Find(map[string]string{})
	if err != nil {
		panic(err)
	}

	result := struct {
		Size int
		Data []services.Entity
	}{len(products), products}
	c.IndentedJSON(http.StatusOK, result)

}

func GetOneProduct(c *gin.Context) {

	id := c.Param("id")

	services.SetDB(DB)

	p := services.Entity.New(&product.Product{})

	p, err := p.FindOne(map[string]string{"Id": id})
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, p)

}

func AddProduct(c *gin.Context) {

	services.SetDB(DB)

	var product *product.Product
	if err := c.BindJSON(&product); err != nil {
		panic(err)
	}

	if _, err := product.Create(); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, product)
}

func DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	services.SetDB(DB)

	p := services.Entity.New(&product.Product{})

	p, err := p.FindOne(map[string]string{"Id": id})
	if err != nil {
		panic(err)
	}

	result, err := p.Delete(map[string]string{"Id": id})
	if err != nil {
		panic(err)
	}

	if !result {
		panic("not deleted")
	}

	c.IndentedJSON(http.StatusOK, p)
}

func PutProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	services.SetDB(DB)
	temp := product.Product{}
	_, err = temp.FindOne(map[string]string{"Id": fmt.Sprint(id)})
	if err != nil {
		panic(err)
	}

	var product product.Product
	product.Id = id
	if err := c.BindJSON(&product); err != nil {
		panic(err)
	}
	p := services.Entity.New(&product)

	log.Println(p)

	p, err = p.Update()
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, p)

}
