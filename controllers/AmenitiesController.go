package controllers

import "github.com/gin-gonic/gin"


func CreateAmenityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateAmenityHandler",
	})
}

func GetAllAmenitiesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAllAmenitiesHandler",
	})
}

func GetHotelAmenitiesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetHotelAmenitiesHandler",
	})
}

func GetAmenityDetailsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAmenityDetailsHandler",
	})
}

func GetFacilitiesForHotelHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetFacilitiesForHotelHandler",
	})
}
