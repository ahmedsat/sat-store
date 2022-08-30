package routes

import (
	"github.com/ahmedsat/sat-store/controllers"
	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	v1:=r.Group("/api/v1")
	v1.GET("/products", controllers.GetAllProducts)
	v1.GET("/products/:id", controllers.GetOneProduct)
	v1.POST("/products", controllers.AddProduct)
}
// services