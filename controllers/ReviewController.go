package controllers

import "github.com/gin-gonic/gin"

func CreateReviewHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateReviewHandler",
	})
}

func GetAllReviewsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAllReviewsHandler",
	})
}

func GetHotelReviewsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetHotelReviewsHandler",
	})
}

func GetReviewDetailsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetReviewDetailsHandler",
	})
}

func UpdateReviewHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateReviewHandler",
	})
}

func DeleteReviewHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteReviewHandler",
	})
}
