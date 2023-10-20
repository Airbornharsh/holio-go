package controllers

import "github.com/gin-gonic/gin"

func CreateUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateUserHandler",
	})
}

func GetUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetUserHandler",
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

func ChangeUserNameHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeUserFirstNameHandler",
	})
}

func ChangeUserFullNameHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeUserFullNameHandler",
	})
}

func ChangeUserAddressHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeUserAddressHandler",
	})
}

func ChangeUserTypeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeUserTypeHandler",
	})
}
