package utils

import (
	"errors"

	"github.com/ahmedsat/sat-store/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) (user models.User, err error) {
	// get user from request in type any
	u, ok := c.Get("user")
	if !ok {
		err = errors.New("you need to login first")
		return
	}
	user = u.(models.User) // change the type from any to models.User
	return
}
