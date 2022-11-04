package main

import (
	"os"

	"github.com/ahmedsat/sat-store/database"
	"github.com/ahmedsat/sat-store/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload" // load .env file
)

func main() {
	AutoMigrate := os.Getenv("AUTO_MIGRATE")
	addr := "localhost:" + os.Getenv("PORT")

	r := gin.Default() // default gin engin

	database.Conn()

	if AutoMigrate == "TRUE" {
		database.MigrateAll()
	}

	routes.RootRoutes(r)
	r.Run(addr)

}
