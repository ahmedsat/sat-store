package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Get Current User")
}
