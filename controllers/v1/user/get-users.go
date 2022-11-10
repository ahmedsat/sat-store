package userController

import (
	"net/http"

	"github.com/ahmedsat/sat-store/database"
	"github.com/ahmedsat/sat-store/models"
	"github.com/ahmedsat/sat-store/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {

	// get search data
	searchQuery := models.UserSearchQueries{}
	c.Bind(&searchQuery)

	// turn search query map to string
	searchConditions, searchValues := models.UserSearchQueriesParser(searchQuery)

	// get user that was provided by auth middleware
	loggedUser, err := utils.GetUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "you need to login first",
		})
		c.Abort()
		return
	}

	// if user have no Privileges return error message
	if loggedUser.Privileges != string(models.ADMEN) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "you are not authorized to do this,admin only",
		})
		c.Abort()
		return
	}

	users := []models.User{}
	var result *gorm.DB

	if len(searchValues) > 0 {

		// convert searchValues from []string []interface{}
		newSearchValues := make([]interface{}, len(searchValues))
		for i, v := range searchValues {
			newSearchValues[i] = v
		}

		result = database.Instance.Where(searchConditions, newSearchValues...).Find(&users)
	} else {
		result = database.Instance.Find(&users)
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		// "q":     searchConditions,
		// "v":     searchValues,
		"count": len(users),
		"users": users,
		// "x":     searchQuery,
	})

}
