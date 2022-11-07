package main

import (
	"os"

	"github.com/ahmedsat/sat-store/database"
	"github.com/ahmedsat/sat-store/middlewares"
	"github.com/ahmedsat/sat-store/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload" // load .env file
)

func main() {
	AutoMigrate := os.Getenv("AUTO_MIGRATE")
	addr := "localhost:" + os.Getenv("PORT")

	database.Conn()

	r := gin.Default() // default gin engin

	r.Use(middlewares.Logger())

	if AutoMigrate == "TRUE" {
		database.MigrateAll()
	}

	routes.RootRoutes(r)
	r.Run(addr)

}
