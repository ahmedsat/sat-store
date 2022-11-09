package userController

import (
	"net/http"
	"time"

	"github.com/ahmedsat/sat-store/database"
	"github.com/ahmedsat/sat-store/models"
	"github.com/ahmedsat/sat-store/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type searchQueries struct {
	ID         uint           `gorm:"primarykey" form:"id"`
	CreatedAt  time.Time      `form:"created_at"`
	UpdatedAt  time.Time      `form:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Name       string         `form:"name" gorm:"default:no name"`
	Username   string         `form:"username" gorm:"unique"`
	Email      string         `form:"email" gorm:"unique"`
	Password   string         `form:"password"`
	Phone      string         `form:"Phone"`
	Address    string         `form:"address"`
	Privileges string         `form:"privileges"`
}

func GetUsers(c *gin.Context) {

	// get search data
	s := searchQueries{}
	c.Bind(&s)

	// tests

	println(s.ID)

	// end tests

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
