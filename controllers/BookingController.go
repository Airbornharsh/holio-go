package controllers

import (
	"fmt"

	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/helpers"
	"github.com/airbornharsh/holio-go/models"
	"github.com/gin-gonic/gin"
)

func AdminGetUserBookingsHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(400, gin.H{
			"message": "You are not authorized to get the bookings of a user",
		})
		return
	}

	userId := c.Param("userId")

	if userId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid User Id",
		})
		return
	}

	DB, _ := database.GetDB()

	query := `SELECT b.booking_id, b.user_id, b.room_id, b.check_in_date, b.check_out_date, b.total_price, b.booking_status FROM bookings b 
	INNER JOIN rooms r ON b.room_id = r.room_id 
	INNER JOIN hotels h ON r.hotel_id = h.hotel_id 
	WHERE b.user_id = '` + userId + `' AND h.owner_user_id = '` + fmt.Sprintf("%d", tempUser.(models.User).UserID) + `';`

	var bookings []models.Booking

	rows, err := DB.Query(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	for rows.Next() {
		var booking models.Booking

		err = rows.Scan(&booking.BookingID, &booking.UserID, &booking.RoomID, &booking.CheckInDate, &booking.CheckOutDate, &booking.TotalPrice, &booking.BookingStatus)

		if helpers.ErrorResponse(c, err) {
			return
		}

		bookings = append(bookings, booking)
	}

	c.JSON(200, gin.H{
		"message":  "User Bookings",
		"bookings": bookings,
	})
}

func AdminCancelBookingHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(400, gin.H{
			"message": "You are not authorized to cancel a booking",
		})
		return
	}

	bookingId := c.Param("bookingId")

	if bookingId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Booking Id",
		})
		return
	}

	DB, _ := database.GetDB()

	var bookingExists bool

	query := "SELECT EXISTS (SELECT 1 FROM bookings WHERE booking_id = '" + bookingId + "');"

	err := DB.QueryRow(query).Scan(&bookingExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !bookingExists {
		c.JSON(400, gin.H{
			"message": "Booking does not exist",
		})
		return
	}

	var hotelExists bool

	query = "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = (SELECT hotel_id FROM rooms WHERE room_id = (SELECT room_id FROM bookings WHERE booking_id = '" + bookingId + "')) AND owner_user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	err = DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel does not exist or you are not authorized to cancel a booking for this hotel",
		})
		return
	}

	query = "UPDATE bookings SET booking_status = 'cancelled' WHERE booking_id = '" + bookingId + "';"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Admin Cancelled the Booking",
	})
}

func AdminConfirmBookingHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(400, gin.H{
			"message": "You are not authorized to confirm a booking",
		})
		return
	}

	bookingId := c.Param("bookingId")

	if bookingId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Booking Id",
		})
		return
	}

	DB, _ := database.GetDB()

	var bookingExists bool

	query := "SELECT EXISTS (SELECT 1 FROM bookings WHERE booking_id = '" + bookingId + "' AND booking_status = 'pending' );"

	err := DB.QueryRow(query).Scan(&bookingExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !bookingExists {
		c.JSON(400, gin.H{
			"message": "Booking does not exist or Booking is Cancelled or Confirmed",
		})
		return
	}

	var hotelExists bool

	query = "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = (SELECT hotel_id FROM rooms WHERE room_id = (SELECT room_id FROM bookings WHERE booking_id = '" + bookingId + "')) AND owner_user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	err = DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel does not exist or you are not authorized to confirm a booking for this hotel",
		})
		return
	}

	query = "UPDATE bookings SET booking_status = 'confirmed' WHERE booking_id = '" + bookingId + "';"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Booking Confirmed",
	})
}

func AdminCheckInHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(400, gin.H{
			"message": "You are not authorized to check in a user",
		})
		return
	}

	bookingId := c.Param("bookingId")

	if bookingId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Booking Id",
		})
		return
	}

	DB, _ := database.GetDB()

	var bookingExists bool

	query := "SELECT EXISTS (SELECT 1 FROM bookings WHERE booking_id = '" + bookingId + "' AND booking_status = 'confirmed' );"

	err := DB.QueryRow(query).Scan(&bookingExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !bookingExists {
		c.JSON(400, gin.H{
			"message": "Booking does not exist or Booking is Cancelled or Checked In",
		})
		return
	}

	var hotelExists bool

	query = "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = (SELECT hotel_id FROM rooms WHERE room_id = (SELECT room_id FROM bookings WHERE booking_id = '" + bookingId + "')) AND owner_user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	err = DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel does not exist or you are not authorized to check in a user for this hotel",
		})
		return
	}

	query = "UPDATE bookings SET booking_status = 'checked_in' WHERE booking_id = '" + bookingId + "';"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Checked In",
	})
}

func AdminCheckOutHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser != nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(400, gin.H{
			"message": "You are not authorized to check out a user",
		})
		return
	}

	bookingId := c.Param("bookingId")

	if bookingId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Booking Id",
		})
		return
	}

	DB, _ := database.GetDB()

	var bookingExists bool

	query := "SELECT EXISTS (SELECT 1 FROM bookings WHERE booking_id = '" + bookingId + "' AND booking_status = 'checked_in' );"

	err := DB.QueryRow(query).Scan(&bookingExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !bookingExists {
		c.JSON(400, gin.H{
			"message": "Booking does not exist or Booking is Cancelled or Checked Out",
		})
		return
	}

	var hotelExists bool

	query = "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = (SELECT hotel_id FROM rooms WHERE room_id = (SELECT room_id FROM bookings WHERE booking_id = '" + bookingId + "')) AND owner_user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	err = DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel does not exist or you are not authorized to check out a user for this hotel",
		})
		return
	}

	query = "UPDATE bookings SET booking_status = 'checked_out' WHERE booking_id = '" + bookingId + "';"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Checked Out",
	})
}
