package controllers

import (
	"fmt"

	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/helpers"
	"github.com/airbornharsh/holio-go/models"
	"github.com/gin-gonic/gin"
)

func CreateHotelHandler(c *gin.Context) {
	tempuser, exists := c.Get("user")

	if !exists && tempuser.(models.User).UserType != "owner" {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var hotel models.Hotel

	c.BindJSON(&hotel)

	query := `INSERT INTO hotels (owner_user_id, name, description, address, phone_number, website_url, email, latitude, longitude, star_rating, avg_rating, avg_price) VALUES ('` + fmt.Sprintf("%d", tempuser.(models.User).UserID)	 + `', '` + hotel.Name + `', '` + hotel.Description + `', '` + hotel.Address + `', '` + hotel.PhoneNumber + `', '` + hotel.WebsiteURL + `', '` + hotel.Email + `', '` + fmt.Sprintf("%.8f", hotel.Latitude) + `', '` + fmt.Sprintf("%.8f", hotel.Longitude) + `', '` + fmt.Sprintf("%.1f", hotel.StarRating) + `', '` + fmt.Sprintf("%.2f", hotel.AvgRating) + `', '` + fmt.Sprintf("%.2f", hotel.AvgPrice) + `');`

	DB, _ := database.GetDB()
	_, err := DB.Exec(query)

	if helpers.ErrorResponse(c, err) {

		return
	}

	c.JSON(200, gin.H{
		"message": "Hotel Created SuccessFully",
	})
}

func SearchHotelsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "SearchHotelsHandler",
	})
}

func GetPopularHotelsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetPopularHotelsHandler",
	})
}

func GetHotelHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetHotelHandler",
	})
}

func UpdateHotelHandler(c *gin.Context) {
	// var hotels []

	// query := `SELECT * FROM hotels WHERE owner_user_id = '` + string(tempuser.(models.User).UserID) + `';`
	c.JSON(200, gin.H{
		"message": "UpdateHotelHandler",
	})
}

func DeleteHotelHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteHotelHandler",
	})
}

func ChangeHotelNameHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelNameHandler",
	})
}

func ChangeHotelDescriptionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelDescriptionHandler",
	})
}

func ChangeHotelAddressHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelAddressHandler",
	})
}

func ChangeHotelPhoneHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelPhoneHandler",
	})
}

func ChangeHotelWebsiteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelWebsiteHandler",
	})
}

func ChangeHotelEmailHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelEmailHandler",
	})
}

func ChangeHotelStarRatingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelStarRatingHandler",
	})
}

func ChangeHotelAvgRatingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelAvgRatingHandler",
	})
}

func ChangeHotelAvgPriceHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelAvgPriceHandler",
	})
}

func ChangeHotelLocationHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelLocationHandler",
	})
}

func ChangeHotelFacilitiesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelFacilitiesHandler",
	})
}

func ChangeHotelRoomsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelRoomsHandler",
	})
}

func ChangeHotelAmenitiesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelAmenitiesHandler",
	})
}
