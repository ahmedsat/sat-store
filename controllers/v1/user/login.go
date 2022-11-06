package userController

import (
	"net/http"
	"os"

	"github.com/ahmedsat/sat-store/auth"
	"github.com/ahmedsat/sat-store/database"
	"github.com/ahmedsat/sat-store/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	providedData := struct {
		Username string `json:"username" gorm:"unique"`
		Password string `json:"password"`
	}{} // user from http request

	// check if request mach user model and bind it
	if err := c.ShouldBindJSON(&providedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := models.User{} // new empty user

	result := database.Instance.Where("username = ?", providedData.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, result.Error.Error())
	}

	// check the Password
	if err := user.CheckPassword(providedData.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// generate JWT
	token, err := auth.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
		"os":    os.Getenv("token"),
	})
}
