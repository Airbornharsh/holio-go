package controllers

import (
	"fmt"

	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/helpers"
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
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "model") {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
	}

	type NewUser struct {
		Username string `json:"username"`
		Address  string `json:"address"`
		FullName string `json:"full_name"`
		UserType string `json:"user_type"`
	}

	var newUser NewUser

	err := c.ShouldBindJSON(&newUser)
	if helpers.ErrorResponse(c, err) {
		return
	}

	query := "UPDATE users SET username = '" + newUser.Username + "', address = '" + newUser.Address + "', full_name = '" + newUser.FullName + "', user_type = '" + newUser.UserType + "' WHERE user_id = " + fmt.Sprintf("%d", tempUser.(models.User).UserID) + ";"

	DB, _ := database.GetDB()
	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Updated the User",
	})
}

func DeleteUserHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "model") {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
	}

	type NewUser struct {
		Username string `json:"username"`
		Address  string `json:"address"`
		FullName string `json:"full_name"`
		UserType string `json:"user_type"`
	}

	var newUser NewUser

	err := c.ShouldBindJSON(&newUser)
	if helpers.ErrorResponse(c, err) {
		return
	}

	query := "DELETE FROM users WHERE user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "';"

	DB, _ := database.GetDB()
	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "User Deleted",
	})
}
