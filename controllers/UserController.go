package controllers

import (
	"github.com/airbornharsh/holio-go/models"
	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "model") {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
	}

	user := tempUser.(models.User)
	user.Password = ""

	c.JSON(200, gin.H{
		"message":  "Got the User",
		"userData": user,
	})
}

func UpdateUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateUserHandler",
	})
}

func DeleteUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteUserHandler",
	})
}
