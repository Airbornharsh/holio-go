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

func ChangeReviewRatingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeReviewRatingHandler",
	})
}

func ChangeReviewTitleHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeReviewTitleHandler",
	})
}

func ChangeReviewDescriptionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeReviewDescriptionHandler",
	})
}

func ChangeReviewUserIdHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeReviewUserIdHandler",
	})
}

func ChangeReviewHotelIdHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeReviewHotelIdHandler",
	})
}
