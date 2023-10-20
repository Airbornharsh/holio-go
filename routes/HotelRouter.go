package routes

import (
	"github.com/airbornharsh/holio-go/controllers"
	"github.com/gin-gonic/gin"
)

func HotelRoutes(r *gin.Engine) {
	//Hotel CRUD
	r.POST("/hotels", controllers.CreateHotelHandler)
	r.GET("/hotels/search", controllers.SearchHotelsHandler)
	r.GET("/hotels/popular", controllers.GetPopularHotelsHandler)
	r.GET("/hotels/:id", controllers.GetHotelHandler)
	r.PUT("/hotels/:id", controllers.UpdateHotelHandler)
	r.DELETE("/hotels/:id", controllers.DeleteHotelHandler)

	//Hotel Details
	r.PUT("/hotels/:id/change-name", controllers.ChangeHotelNameHandler)
	r.PUT("/hotels/:id/change-description", controllers.ChangeHotelDescriptionHandler)
	r.PUT("/hotels/:id/change-address", controllers.ChangeHotelAddressHandler)
	r.PUT("/hotels/:id/change-phone", controllers.ChangeHotelPhoneHandler)
	r.PUT("/hotels/:id/change-website", controllers.ChangeHotelWebsiteHandler)
	r.PUT("/hotels/:id/change-email", controllers.ChangeHotelEmailHandler)
	r.PUT("/hotels/:id/change-star-rating", controllers.ChangeHotelStarRatingHandler)
	r.PUT("/hotels/:id/change-avg-rating", controllers.ChangeHotelAvgRatingHandler)
	r.PUT("/hotels/:id/change-avg-price", controllers.ChangeHotelAvgPriceHandler)
	r.PUT("/hotels/:id/change-location", controllers.ChangeHotelLocationHandler)
	r.PUT("/hotels/:id/change-facilities", controllers.ChangeHotelFacilitiesHandler)
	r.PUT("/hotels/:id/change-rooms", controllers.ChangeHotelRoomsHandler)
	r.PUT("/hotels/:id/change-amenities", controllers.ChangeHotelAmenitiesHandler)
}
