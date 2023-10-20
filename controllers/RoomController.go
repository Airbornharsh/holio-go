package controllers

import "github.com/gin-gonic/gin"

func CreateRoomHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateRoomHandler",
	})
}

func SearchRoomsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "SearchRoomsHandler",
	})
}

func GetPopularRoomsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetPopularRoomsHandler",
	})
}

func GetRoomHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetRoomHandler",
	})
}

func UpdateRoomHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateRoomHandler",
	})
}

func DeleteRoomHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteRoomHandler",
	})
}

func GetRoomsForHotelHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetRoomsForHotelHandler",
	})
}

func GetRoomDetailsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetRoomDetailsHandler",
	})
}

func CheckRoomAvailabilityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CheckRoomAvailabilityHandler",
	})
}

func CreateBookingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateBookingHandler",
	})
}

func GetUserBookingsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetUserBookingsHandler",
	})
}

func CancelBookingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CancelBookingHandler",
	})
}

func ChangeRoomNameHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomNameHandler",
	})
}

func ChangeRoomDescriptionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomDescriptionHandler",
	})
}

func ChangeRoomPriceHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomPriceHandler",
	})
}

func ChangeRoomCapacityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomCapacityHandler",
	})
}

func ChangeRoomPhotosHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomImagesHandler",
	})
}

func ChangeRoomAvailabilityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomStatusHandler",
	})
}

func ChangeRoomAvgRatingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomAvgRatingHandler",
	})
}

func ChangeRoomAvgPriceHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomAvgPriceHandler",
	})
}

func ChangeRoomLocationHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomLocationHandler",
	})
}

func ChangeRoomFacilitiesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomFacilitiesHandler",
	})
}

func ChangeRoomAmenitiesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomAmenitiesHandler",
	})
}
