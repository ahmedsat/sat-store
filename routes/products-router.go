package routes

import (
	"github.com/ahmedsat/sat-store/controllers"
	"github.com/gin-gonic/gin"
)

func productRoutes(rg *gin.RouterGroup) {
	product := rg.Group("/products")

	product.POST("/", controllers.AddProduct)
	product.GET("/", controllers.GetAllProducts)
	product.GET("/:id", controllers.GetOneProduct)
	product.PUT("/:id", controllers.PutProduct)
	product.DELETE("/:id", controllers.DeleteProduct)
}
