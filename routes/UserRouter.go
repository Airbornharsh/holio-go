package routes

import (
	// "github.com/airbornharsh/holio-go/controllers"
	"github.com/airbornharsh/holio-go/controllers"
	"github.com/airbornharsh/holio-go/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/user/:id", middlewares.TokenVerifyMiddleWare, controllers.GetUserHandler)
	r.PUT("/user", middlewares.TokenVerifyMiddleWare, controllers.UpdateUserHandler)
	r.DELETE("/user", middlewares.TokenVerifyMiddleWare, controllers.DeleteUserHandler)
}
