package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {
	v, ok := c.Get("userID")
	c.JSON(http.StatusOK, gin.H{
		"status": ok,
		"value":  v,
	})
}
