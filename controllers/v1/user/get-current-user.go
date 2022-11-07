package userController

import (
	"net/http"

	"github.com/ahmedsat/sat-store/database"
	"github.com/ahmedsat/sat-store/models"
	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "you net to login first",
		})
	}

	user := models.User{}

	database.Instance.First(&user, id)

	c.JSON(http.StatusOK, gin.H{
		"name":     user.Name,
		"email":    user.Email,
		"username": user.Username,
	})
}
