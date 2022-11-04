package main

import (
	"github.com/ahmedsat/sat-store/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload" // load .env file
)

func main() {

	r := gin.Default() // default gin engin
	routes.RootRoutes(r)
	r.Run("localhost:8080")

}
