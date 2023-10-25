package routes

import (
	"github.com/airbornharsh/holio-go/controllers"
	"github.com/airbornharsh/holio-go/middlewares"
	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine) {
	r.POST("/bookings", middlewares.TokenVerifyMiddleWare, controllers.CreateBookingHandler)
	r.GET("/users/bookings", middlewares.TokenVerifyMiddleWare, controllers.GetUserBookingsHandler)
	r.DELETE("/users/bookings/:bookingId", middlewares.TokenVerifyMiddleWare, controllers.CancelBookingHandler)
	r.GET("/admin/users/:userId/bookings", middlewares.TokenVerifyMiddleWare, controllers.AdminGetUserBookingsHandler)
	r.PUT("/admin/bookings/:bookingId", middlewares.TokenVerifyMiddleWare, controllers.AdminConfirmBookingHandler)
	r.DELETE("/admin/bookings/:bookingId", middlewares.TokenVerifyMiddleWare, controllers.AdminCancelBookingHandler)
	r.PUT("/admin/bookings/:bookingId/checkin", middlewares.TokenVerifyMiddleWare, controllers.AdminCheckInHandler)
	r.PUT("/admin/bookings/:bookingId/checkout", middlewares.TokenVerifyMiddleWare, controllers.AdminCheckOutHandler)

}
