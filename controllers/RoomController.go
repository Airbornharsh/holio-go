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

	DB, err := database.GetDB()

	if helpers.ErrorResponse(c, err) {
		return
	}

	err = DB.QueryRow(query).Scan(&room.RoomID, &room.HotelID, &room.RoomNumber, &room.RoomType, &room.RoomCapacity, &room.Description, &room.Price, &room.Rating, &room.Availability)

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

	DB, err := database.GetDB()

	if helpers.ErrorResponse(c, err) {
		return
	}

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

func ChangeRoomAvailabilityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeRoomStatusHandler",
	})
}
