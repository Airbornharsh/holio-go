package routes

import (
	"github.com/gin-gonic/gin"
)

// Routes is the main function that will be called from main.go
func Router(r *gin.Engine) {
	AuthRoutes(r)
	UserRoutes(r)
	HotelRoutes(r)
	RoomRoutes(r)
	ReviewsRoutes(r)
	AmenitiesRoutes(r)
}
