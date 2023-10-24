package controllers

import (
	"fmt"

	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/helpers"
	"github.com/airbornharsh/holio-go/models"
	"github.com/gin-gonic/gin"
)

func CreateRoomHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(400, gin.H{
			"message": "You are not authorized to create a room",
		})
		return
	}

	var room *models.Room

	err := c.BindJSON(&room)

	if helpers.ErrorResponse(c, err) {
		return
	}

	DB, _ := database.GetDB()

	var hotelExists bool
	query := "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = '" + fmt.Sprintf("%d", room.HotelID) + "' AND owner_user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	err = DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel does not exist or you are not authorized to create a room for this hotel",
		})
		return
	}

	query = "INSERT into rooms (hotel_id,room_number,room_type,room_capacity,description,price,rating,availability) VALUES ('" + fmt.Sprintf("%d", room.HotelID) + "','" + fmt.Sprintf("%d", room.RoomNumber) + "','" + room.RoomType + "','" + fmt.Sprintf("%d", room.RoomCapacity) + "','" + room.Description + "','" + fmt.Sprintf("%.2f", room.Price) + "','" + fmt.Sprintf("%.2f", room.Rating) + "','" + fmt.Sprintf("%t", room.Availability) + "');"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Room Created",
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
	roomId := c.Param("id")

	if roomId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Room Id",
		})
		return
	}

	query := "SELECT * FROM rooms WHERE room_id = '" + roomId + "';"

	var room models.Room

	DB, _ := database.GetDB()
	err := DB.QueryRow(query).Scan(&room.RoomID, &room.HotelID, &room.RoomNumber, &room.RoomType, &room.RoomCapacity, &room.Description, &room.Price, &room.Rating, &room.Availability)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Got the Hotel ",
		"room":    room,
	})
}

func UpdateRoomHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(400, gin.H{
			"message": "You are not authorized to update a room",
		})
		return
	}

	roomId := c.Param("id")

	if roomId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Room Id",
		})
		return
	}

	var newRoom *models.Room

	err := c.BindJSON(&newRoom)

	if helpers.ErrorResponse(c, err) {
		return
	}

	DB, _ := database.GetDB()

	var hotelId int
	query := "SELECT hotel_id FROM rooms WHERE room_id = '" + roomId + "';"

	err = DB.QueryRow(query).Scan(&hotelId)

	if helpers.ErrorResponse(c, err) {
		return
	}

	var hotelExists bool
	query = "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = '" + fmt.Sprintf("%d", hotelId) + "' AND owner_user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	err = DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel does not exist or you are not authorized to update a room for this hotel",
		})
		return
	}

	var room models.Room

	query = "SELECT * FROM rooms WHERE room_id = '" + roomId + "';"

	err = DB.QueryRow(query).Scan(&room.RoomID, &room.HotelID, &room.RoomNumber, &room.RoomType, &room.RoomCapacity, &room.Description, &room.Price, &room.Rating, &room.Availability)

	if helpers.ErrorResponse(c, err) {
		return
	}

	helpers.ReplaceRoom(&room, newRoom)

	query = "UPDATE rooms SET room_number = '" + fmt.Sprintf("%d", room.RoomNumber) + "', room_type = '" + room.RoomType + "', room_capacity = '" + fmt.Sprintf("%d", room.RoomCapacity) + "', description = '" + room.Description + "', price = '" + fmt.Sprintf("%.2f", room.Price) + "', rating = '" + fmt.Sprintf("%.2f", room.Rating) + "', availability = '" + fmt.Sprintf("%t", room.Availability) + "' WHERE room_id = '" + roomId + "';"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Room Updated",
	})
}

func DeleteRoomHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(400, gin.H{
			"message": "You are not authorized to delete a room",
		})
		return
	}

	roomId := c.Param("id")

	if roomId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Room Id",
		})
		return
	}

	var hotelId int
	DB, _ := database.GetDB()

	query := "SELECT hotel_id FROM rooms WHERE room_id = '" + roomId + "';"

	err := DB.QueryRow(query).Scan(&hotelId)

	if helpers.ErrorResponse(c, err) {
		return
	}

	var hotelExists bool
	query = "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = '" + fmt.Sprintf("%d", hotelId) + "' AND owner_user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	err = DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel does not exist or you are not authorized to delete a room for this hotel",
		})
		return
	}

	query = "DELETE FROM rooms WHERE room_id = '" + roomId + "';"

	if helpers.ErrorResponse(c, err) {
		return
	}

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Room Deleted Successfully",
	})
}

func GetRoomsForHotelHandler(c *gin.Context) {
	hotelId := c.Param("hotelId")

	if hotelId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Hotel Id",
		})
		return
	}

	query := "SELECT * FROM rooms WHERE hotel_id = '" + hotelId + "';"

	var rooms []models.Room

	DB, _ := database.GetDB()

	rows, err := DB.Query(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	for rows.Next() {
		var room models.Room

		err = rows.Scan(&room.RoomID, &room.HotelID, &room.RoomNumber, &room.RoomType, &room.RoomCapacity, &room.Description, &room.Price, &room.Rating, &room.Availability)

		if helpers.ErrorResponse(c, err) {
			return
		}

		rooms = append(rooms, room)
	}

	if rooms == nil {
		c.JSON(400, gin.H{
			"message": "No Rooms Found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Rooms Found",
		"rooms":   rooms,
	})
}

func CreateBookingHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil) {
		c.JSON(400, gin.H{
			"message": "You are not authorized to create a booking",
		})
		return
	}

	var booking *models.Booking

	err := c.BindJSON(&booking)

	if helpers.ErrorResponse(c, err) {
		return
	}

	DB, _ := database.GetDB()

	var roomExists bool

	query := "SELECT EXISTS (SELECT 1 FROM rooms WHERE room_id = '" + fmt.Sprintf("%d", booking.RoomID) + "');"

	err = DB.QueryRow(query).Scan(&roomExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !roomExists {
		c.JSON(400, gin.H{
			"message": "Room does not exist",
		})
		return
	}

	var roomAvailability bool

	query = "SELECT availability FROM rooms WHERE room_id = '" + fmt.Sprintf("%d", booking.RoomID) + "';"

	err = DB.QueryRow(query).Scan(&roomAvailability)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !roomAvailability {
		c.JSON(400, gin.H{
			"message": "Room is not available",
		})
		return
	}

	var hotelId int

	query = "SELECT hotel_id FROM rooms WHERE room_id = '" + fmt.Sprintf("%d", booking.RoomID) + "';"

	err = DB.QueryRow(query).Scan(&hotelId)

	if helpers.ErrorResponse(c, err) {
		return
	}

	var hotelExists bool

	query = "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = '" + fmt.Sprintf("%d", hotelId) + "');"

	err = DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel does not exist",
		})
		return
	}

	query = "INSERT into bookings (user_id,room_id,check_in_date,check_out_date,total_price,booking_status) VALUES ('" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "','" + fmt.Sprintf("%d", booking.RoomID) + "','" + booking.CheckInDate + "','" + booking.CheckOutDate + "','" + fmt.Sprintf("%.2f", booking.TotalPrice) + "','" + booking.BookingStatus + "') RETURNING booking_id;"

	var bookingID int
	row, err := DB.Query(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	defer row.Close()

	row.Next()

	err = row.Scan(&bookingID)

	if helpers.ErrorResponse(c, err) {
		return
	}

	booking.UserID = tempUser.(models.User).UserID
	booking.BookingID = bookingID

	c.JSON(200, gin.H{
		"message": "Booking Created Successfully",
		"booking": booking,
	})
}

func GetUserBookingsHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser != nil) {
		c.JSON(401, gin.H{
			"message": "No User",
		})
	}

	DB, _ := database.GetDB()

	query := "SELECT * FROM bookings WHERE user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "';"

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
		"message":  "Got Bookings",
		"bookings": bookings,
	})
}

func CancelBookingHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil) {
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

	query := "SELECT EXISTS (SELECT 1 FROM bookings WHERE booking_id = '" + bookingId + "' AND user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	err := DB.QueryRow(query).Scan(&bookingExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !bookingExists {
		c.JSON(400, gin.H{
			"message": "Booking does not exist or you are not authorized to cancel this booking",
		})
		return
	}

	query = "UPDATE bookings SET booking_status = 'cancelled' WHERE booking_id = '" + bookingId + "';"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Cancelled the Booking",
	})
}

func ChangeRoomAvailabilityHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(400, gin.H{
			"message": "You are not authorized to change the availability of a room",
		})
		return
	}

	roomId := c.Param("id")

	if roomId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Room Id",
		})
		return
	}

	type RoomAvailability struct {
		Availability bool `json:"availability"`
	}

	var availability *RoomAvailability

	err := c.BindJSON(&availability)

	if helpers.ErrorResponse(c, err) {
		return
	}

	DB, _ := database.GetDB()

	var hotelId int

	query := "SELECT hotel_id FROM rooms WHERE room_id = '" + roomId + "';"

	err = DB.QueryRow(query).Scan(&hotelId)

	if helpers.ErrorResponse(c, err) {
		return
	}

	var hotelExists bool

	query = "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = '" + fmt.Sprintf("%d", hotelId) + "' AND owner_user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	err = DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel does not exist or you are not authorized to change the availability of a room for this hotel",
		})
		return
	}

	query = "UPDATE rooms SET availability = '" + fmt.Sprintf("%t", availability.Availability) + "' WHERE room_id = '" + roomId + "';"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Changed Room Availability",
	})
}
