package controllers

import "github.com/gin-gonic/gin"

func SignupHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "SignupHandler",
	})
}

func LoginHanlder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "LoginHanlder",
	})
}

func LogoutHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "LogoutHandler",
	})
}

func ForgotPasswordHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ForgotPasswordHandler",
	})
}

func ResetPasswordHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ResetPasswordHandler",
	})
}

func ChangePasswordHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangePasswordHandler",
	})
}

func ChangeEmailHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeEmailHandler",
	})
}

func ChangePhoneHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangePhoneHandler",
	})
}
