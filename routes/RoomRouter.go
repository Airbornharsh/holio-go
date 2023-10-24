package routes

import (
	"github.com/airbornharsh/holio-go/controllers"
	"github.com/airbornharsh/holio-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RoomRoutes(r *gin.Engine) {
	// Rooms
	r.POST("/rooms", middlewares.TokenVerifyMiddleWare, controllers.CreateRoomHandler)
	// r.GET("/rooms/search", controllers.SearchRoomsHandler)
	// r.GET("/rooms/popular", controllers.GetPopularRoomsHandler)
	// r.GET("/rooms/:id", controllers.GetRoomHandler)
	// r.PUT("/rooms/:id", controllers.UpdateRoomHandler)
	// r.DELETE("/rooms/:id", controllers.DeleteRoomHandler)

	//Room Details
	// r.GET("/hotels/:hotelId/rooms", controllers.GetRoomsForHotelHandler)
	// r.GET("/hotels/:hotelId/rooms/:roomId", controllers.GetRoomDetailsHandler)

	//Bookinga
	// r.GET("/hotels/:hotelId/rooms/availability", controllers.CheckRoomAvailabilityHandler)
	// r.POST("/bookings", controllers.CreateBookingHandler)
	// r.GET("/users/:userId/bookings", controllers.GetUserBookingsHandler)
	// r.DELETE("/users/:userId/bookings/:bookingId", controllers.CancelBookingHandler)

	//Changing the Availability of a Room
	// r.PUT("/rooms/:id/change-availability", controllers.ChangeRoomAvailabilityHandler)
}
