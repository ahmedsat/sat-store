package utils

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func LogFile(filePath string, consol bool) {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create(filePath)
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	if consol {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
}
