package routes

import (
	"github.com/gin-gonic/gin"
)

func RootRoutes(r *gin.Engine) {

	// r.Static("/assets", "./assets")
	// r.StaticFS("/more_static", http.Dir("my_file_system"))
	// r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	V1(r.Group("/api/v1"))

}
