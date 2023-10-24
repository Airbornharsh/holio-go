package routes

import (
	"github.com/airbornharsh/holio-go/controllers"
	"github.com/airbornharsh/holio-go/middlewares"
	"github.com/gin-gonic/gin"
)

func HotelRoutes(r *gin.Engine) {
	//Anyone
	// query - hotel_name=Ar & price_start=0 & price_end=4000
	r.GET("/hotels/search", controllers.SearchHotelsHandler)
	// query - min_star=4
	r.GET("/hotels/popular", controllers.GetPopularHotelsHandler)
	r.GET("/hotel/:id", controllers.GetHotelHandler)

	//Owner
	//Hotel CRUD
	r.POST("/hotel", middlewares.TokenVerifyMiddleWare, controllers.CreateHotelHandler)
	r.PUT("/hotel/:id", middlewares.TokenVerifyMiddleWare, controllers.UpdateHotelHandler)
	r.DELETE("/hotel/:id", middlewares.TokenVerifyMiddleWare, controllers.DeleteHotelHandler)

	//Hotel Details
	// r.PUT("/hotels/:id/change-facilities", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelFacilitiesHandler)
	// r.PUT("/hotels/:id/change-rooms", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelRoomsHandler)
	// r.PUT("/hotels/:id/change-amenities", middlewares.TokenVerifyMiddleWare, controllers.ChangeHotelAmenitiesHandler)
}
