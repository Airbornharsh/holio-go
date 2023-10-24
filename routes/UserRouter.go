package routes

import (
	// "github.com/airbornharsh/holio-go/controllers"
	"github.com/airbornharsh/holio-go/controllers"
	"github.com/airbornharsh/holio-go/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/user/:id", middlewares.TokenVerifyMiddleWare, controllers.GetUserHandler)
	// r.PUT("/user/:id", controllers.UpdateUserHandler)
	// r.DELETE("/user/:id", controllers.DeleteUserHandler)
}
