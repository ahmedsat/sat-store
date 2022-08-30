package routes

import (
	"github.com/ahmedsat/sat-store/controllers"
	"github.com/gin-gonic/gin"
)

func productRoutes(rg *gin.RouterGroup){
	product:=rg.Group("/products")

	product.GET("/", controllers.GetAllProducts)
	product.GET("/:id", controllers.GetOneProduct)
	product.POST("/", controllers.AddProduct)
	product.DELETE("/:id",controllers.DeleteProduct)
}