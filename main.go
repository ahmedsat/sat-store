package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload" // load .env file
)

func main() {

	r := gin.Default() // default gin engin

	r.Run("localhost:8080")

}
