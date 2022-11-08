package routes

import (
	userController "github.com/ahmedsat/sat-store/controllers/v1/user"
	"github.com/ahmedsat/sat-store/middlewares"
	"github.com/gin-gonic/gin"
)

func V1(rg *gin.RouterGroup) {

	userRouts(rg.Group("/user"))
	productRouts(rg.Group("/product"))
	categoryRouter(rg.Group("/category"))
	cartRouter(rg.Group("/cart"))
	orderRouter(rg.Group("/order"))

}

// user router
func userRouts(rg *gin.RouterGroup) {
	rg.GET("/", middlewares.Auth(), userController.GetAllUsers)
	rg.GET("/me", middlewares.Auth(), userController.GetCurrentUser)
	rg.POST("/login", userController.Login)
	rg.POST("/register", userController.Register)

}

// product router
func productRouts(rg *gin.RouterGroup) {
	rg.GET("/", mocController)       // get all && search
	rg.GET("/:id", mocController)    // get one by id
	rg.POST("/", mocController)      // create one
	rg.PATCH("/:id", mocController)  // update fields
	rg.PUT("/:id", mocController)    // replace one
	rg.DELETE("/:id", mocController) // delete one
}

// category router
func categoryRouter(rg *gin.RouterGroup) {
	rg.GET("/", mocController)       // get all && search
	rg.GET("/:id", mocController)    // get one by id
	rg.POST("/", mocController)      // create one
	rg.PATCH("/:id", mocController)  // update fields
	rg.PUT("/:id", mocController)    // replace one
	rg.DELETE("/:id", mocController) // delete one
}

// cart router
func cartRouter(rg *gin.RouterGroup) {
	rg.GET("/", mocController)       // get all && search
	rg.GET("/:id", mocController)    // get one by id
	rg.POST("/", mocController)      // create one
	rg.PATCH("/:id", mocController)  // update fields
	rg.PUT("/:id", mocController)    // replace one
	rg.DELETE("/:id", mocController) // delete one
}

// order router
func orderRouter(rg *gin.RouterGroup) {
	rg.GET("/", mocController)       // get all && search
	rg.GET("/:id", mocController)    // get one by id
	rg.POST("/", mocController)      // create one
	rg.PATCH("/:id", mocController)  // update fields
	rg.PUT("/:id", mocController)    // replace one
	rg.DELETE("/:id", mocController) // delete one
}
