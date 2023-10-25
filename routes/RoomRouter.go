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
	r.GET("/room/:id", controllers.GetRoomHandler)
	r.PUT("/room/:id", middlewares.TokenVerifyMiddleWare, controllers.UpdateRoomHandler)
	r.DELETE("/room/:id", middlewares.TokenVerifyMiddleWare, controllers.DeleteRoomHandler)

	//Room Details
	r.GET("/hotels/:hotelId/rooms", controllers.GetRoomsForHotelHandler)

	//Bookinga
	r.POST("/bookings", middlewares.TokenVerifyMiddleWare, controllers.CreateBookingHandler)
	r.GET("/users/bookings", middlewares.TokenVerifyMiddleWare, controllers.GetUserBookingsHandler)
	r.GET("/admin/users/:userId/bookings", middlewares.TokenVerifyMiddleWare, controllers.GetAdminUserBookingsHandler)
	r.DELETE("/users/bookings/:bookingId", middlewares.TokenVerifyMiddleWare, controllers.CancelBookingHandler)
	r.PUT("/admin/bookings/:bookingId", middlewares.TokenVerifyMiddleWare, controllers.AdminConfirmBookingHandler)
	r.DELETE("/admin/bookings/:bookingId", middlewares.TokenVerifyMiddleWare, controllers.AdminCancelBookingHandler)
	// r.PUT("/users/bookings/:bookingId", middlewares.TokenVerifyMiddleWare, controllers.AdminCheckOutHandler)
	// r.PUT("/users/bookings/:bookingId", middlewares.TokenVerifyMiddleWare, controllers.AdminCheckOutHandler)

	//Changing the Availability of a Room
	r.PUT("/room/:id/change-availability", middlewares.TokenVerifyMiddleWare, controllers.ChangeRoomAvailabilityHandler)
}
