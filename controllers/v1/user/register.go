package userController

import (
	"net/http"

	"github.com/ahmedsat/sat-store/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	user := models.User{}
	c.JSON(http.StatusOK, gin.H{
		"name":     user.Name,
		"username": user.Username,
		"token":    "",
	})
}
