package routes

import (
	"github.com/airbornharsh/holio-go/controllers"
	"github.com/airbornharsh/holio-go/middlewares"
	"github.com/gin-gonic/gin"
)

func ReviewsRoutes(r *gin.Engine) {
	r.POST("/reviews", middlewares.TokenVerifyMiddleWare, controllers.CreateReviewHandler)
	r.GET("/reviews", middlewares.TokenVerifyMiddleWare, controllers.GetAllReviewsHandler)

	//Hotel Reviews
	// r.GET("/hotels/:hotelId/reviews", controllers.GetHotelReviewsHandler)
	// r.GET("/hotels/:hotelId/reviews/:reviewId", controllers.GetReviewDetailsHandler)

	//Review CRUD
	// r.PUT("/reviews/:reviewId", controllers.UpdateReviewHandler)
	// r.DELETE("/reviews/:reviewId", controllers.DeleteReviewHandler)
}
