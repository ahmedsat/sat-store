package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/ahmedsat/sat-store/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenList, ok := c.Request.Header["Authorization"]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "you net to supply an Authorization token",
			})
			c.Abort()
		}

		splitToken := strings.Split(tokenList[0], " ")
		if splitToken[0] != os.Getenv("JWT_BEARER") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "you net to supply an valid token",
			})
			c.Abort()
		}

		claims, err := auth.ValidateToken(splitToken[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": err.Error(),
			})
			c.Abort()
		}
		c.Set("userID", claims.ID)

		c.Next()

	}
}
