package userController

import (
	"net/http"

	"github.com/ahmedsat/sat-store/database"
	"github.com/ahmedsat/sat-store/models"
	"github.com/ahmedsat/sat-store/utils"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	// get user that was provided by auth middleware
	user, err := utils.GetUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "you need to login first",
		})
		c.Abort()
		return
	}

	// if user have no Privileges return error message
	if user.Privileges != string(models.ADMEN) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "you are not authorized to do this,admin only",
		})
		c.Abort()
		return
	}

	users := []models.User{}
	result := database.Instance.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(users),
		"users": users,
	})

}
