package routes

import (
	"github.com/gin-gonic/gin"
)

func AmenitiesRoutes(r *gin.Engine) {
	r.POST("/amenities", controllers.CreateAmenityHandler)
	r.GET("/amenities", controllers.GetAllAmenitiesHandler)
	r.GET("/hotels/:hotelId/amenities", controllers.GetHotelAmenitiesHandler)
	r.GET("/hotels/:hotelId/amenities/:amenityId", controllers.GetAmenityDetailsHandler)
	r.GET("/hotels/:hotelId/facilities", controllers.GetFacilitiesForHotelHandler)
}
