package controllers

import (
	"fmt"
	"time"

	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/helpers"
	"github.com/airbornharsh/holio-go/models"
	"github.com/gin-gonic/gin"
)

func CreateReviewHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil) {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var review models.Review

	err := c.ShouldBindJSON(&review)

	if helpers.ErrorResponse(c, err) {
		return
	}

	review.UserID = tempUser.(models.User).UserID
	review.ReviewDate = time.Now().Format("2006-01-02T15:04:05")

	DB, _ := database.GetDB()

	var bookingExists bool

	query := "SELECT EXISTS (SELECT 1 FROM bookings b INNER JOIN rooms r ON b.room_id = r.room_id WHERE b.user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "' AND r.hotel_id = '" + fmt.Sprintf("%d", review.HotelID) + "');"

	err = DB.QueryRow(query).Scan(&bookingExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !bookingExists {
		c.JSON(400, gin.H{
			"message": "You have not booked this hotel",
		})
		return
	}

	var reviewExists bool

	query = "SELECT EXISTS ( SELECT 1 FROM reviews WHERE user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "' AND hotel_id = '" + fmt.Sprintf("%d", review.HotelID) + "')"

	err = DB.QueryRow(query).Scan(&reviewExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if reviewExists {
		c.JSON(400, gin.H{
			"message": "You have already reviewed this hotel",
		})
		return
	}

	query = "INSERT INTO reviews (user_id, hotel_id, rating, review_text, review_date) VALUES ('" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "', '" + fmt.Sprintf("%d", review.HotelID) + "', '" + fmt.Sprintf("%d", review.Rating) + "', '" + review.ReviewText + "', '" + review.ReviewDate + "')"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	query = "SELECT AVG(rating) FROM reviews WHERE hotel_id = '" + fmt.Sprintf("%d", review.HotelID) + "'"

	var avgRating float64

	err = DB.QueryRow(query).Scan(&avgRating)

	if helpers.ErrorResponse(c, err) {
		return
	}

	query = "UPDATE hotels SET avg_rating = '" + fmt.Sprintf("%f", avgRating) + "' WHERE hotel_id = '" + fmt.Sprintf("%d", review.HotelID) + "'"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

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
