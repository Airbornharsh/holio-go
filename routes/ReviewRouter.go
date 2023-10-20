package routes

import (
	"github.com/gin-gonic/gin"
)

func ReviewsRoutes(r *gin.Engine) {
	r.POST("/reviews", controllers.CreateReviewHandler)
	r.GET("/reviews", controllers.GetAllReviewsHandler)

	//Hotel Reviews
	r.GET("/hotels/:hotelId/reviews", controllers.GetHotelReviewsHandler)
	r.GET("/hotels/:hotelId/reviews/:reviewId", controllers.GetReviewDetailsHandler)

	//Review CRUD
	r.PUT("/reviews/:reviewId", controllers.UpdateReviewHandler)
	r.DELETE("/reviews/:reviewId", controllers.DeleteReviewHandler)
	r.PUT("/reviews/:reviewId/change-rating", controllers.ChangeReviewRatingHandler)
	r.PUT("/reviews/:reviewId/change-title", controllers.ChangeReviewTitleHandler)
	r.PUT("/reviews/:reviewId/change-description", controllers.ChangeReviewDescriptionHandler)
	r.PUT("/reviews/:reviewId/change-user-id", controllers.ChangeReviewUserIdHandler)
	r.PUT("/reviews/:reviewId/change-hotel-id", controllers.ChangeReviewHotelIdHandler)
}
