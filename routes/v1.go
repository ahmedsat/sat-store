package routes

import (
	userController "github.com/ahmedsat/sat-store/controllers/v1/user"
	"github.com/gin-gonic/gin"
)

func V1(rg *gin.RouterGroup) {

	userRouts(rg.Group("/user"))

}

func userRouts(rg *gin.RouterGroup) {

	rg.GET("/", userController.GetCurrentUser)
	rg.POST("/login", userController.Login)
	rg.POST("/register", userController.Register)
}
