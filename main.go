package main

import (
	"os"

	"github.com/ahmedsat/sat-store/database"
	"github.com/ahmedsat/sat-store/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload" // load .env file
)

func main() {

	r := gin.Default() // default gin engin

	database.Conn()

	AutoMigrate := os.Getenv("AUTO_MIGRATE")
	if AutoMigrate == "TRUE" {
		database.MigrateAll()
	}

	routes.RootRoutes(r)
	r.Run("localhost:8080")

}
