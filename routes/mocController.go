package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mocController(c *gin.Context) {

	url:=c.Request.RequestURI
	method:=c.Request.Method

	c.JSON(http.StatusOK, gin.H{
		"url":url,
		"method":method,
	})
}
