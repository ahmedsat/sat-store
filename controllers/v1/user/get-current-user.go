package userController

import (
	"net/http"

	"github.com/ahmedsat/sat-store/utils"
	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {

	// get user that was provided by auth middleware
	user, err := utils.GetUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "you need to login first",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":     user.Name,
		"email":    user.Email,
		"username": user.Username,
	})
}
