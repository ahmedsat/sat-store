package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/ahmedsat/sat-store/auth"
	"github.com/ahmedsat/sat-store/database"
	"github.com/ahmedsat/sat-store/models"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// token was provided in request header as a slice of string with name "Authorization"
		tokenList, ok := c.Request.Header["Authorization"]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "you net to supply an Authorization token",
			})
			c.Abort()
		}

		splitToken := strings.Split(tokenList[0], " ") // first token on the list is needed , and split
		// token have to start with the correct bearer
		if splitToken[0] != os.Getenv("JWT_BEARER") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "you net to supply an valid token",
			})
			c.Abort()
		}

		// validate the token and get claims
		claims, err := auth.ValidateToken(splitToken[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": err.Error(),
			})
			c.Abort()
		}

		// check if the user in database
		user := models.User{}
		result := database.Instance.First(&user, claims.ID)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "you need to login first",
			})
			c.Abort()
			return
		}

		// bind user to the request
		c.Set("user", user)

		c.Next() // go to next handler

	}
}
