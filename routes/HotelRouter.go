package routes

import (
	"github.com/airbornharsh/holio-go/controllers"
	"github.com/airbornharsh/holio-go/middlewares"
	"github.com/gin-gonic/gin"
)

func HotelRoutes(r *gin.Engine) {
	//Hotel CRUD
	r.POST("/hotel", middlewares.TokenVerifyMiddleWare, controllers.CreateHotelHandler)
	r.GET("/hotels/search", controllers.SearchHotelsHandler)
	r.GET("/hotels/popular", controllers.GetPopularHotelsHandler)
	r.GET("/hotel/:id", controllers.GetHotelHandler)
	r.PUT("/hotel/:id", middlewares.TokenVerifyMiddleWare, controllers.UpdateHotelHandler)
	r.DELETE("/hotel/:id", middlewares.TokenVerifyMiddleWare, controllers.DeleteHotelHandler)

	//Hotel Details
	r.PUT("/hotels/:id/change-name", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelNameHandler)
	r.PUT("/hotels/:id/change-description", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelDescriptionHandler)
	r.PUT("/hotels/:id/change-address", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelAddressHandler)
	r.PUT("/hotels/:id/change-phone", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelPhoneHandler)
	r.PUT("/hotels/:id/change-website", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelWebsiteHandler)
	r.PUT("/hotels/:id/change-email", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelEmailHandler)
	r.PUT("/hotels/:id/change-star-rating", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelStarRatingHandler)
	r.PUT("/hotels/:id/change-avg-rating", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelAvgRatingHandler)
	r.PUT("/hotels/:id/change-avg-price", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelAvgPriceHandler)
	r.PUT("/hotels/:id/change-location", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelLocationHandler)
	r.PUT("/hotels/:id/change-facilities", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelFacilitiesHandler)
	r.PUT("/hotels/:id/change-rooms", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelRoomsHandler)
	r.PUT("/hotels/:id/change-amenities", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelAmenitiesHandler)
}
