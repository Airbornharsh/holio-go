package routes

import (
	"github.com/airbornharsh/holio-go/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/user", controllers.CreateUserHandler)
	r.GET("/user/:id", controllers.GetUserHandler)
	r.PUT("/user/:id", controllers.UpdateUserHandler)
	r.DELETE("/user/:id", controllers.DeleteUserHandler)
	r.PUT("/user/:id/change-username", controllers.ChangeUserNameHandler)
	r.PUT("/user/:id/change-fullname", controllers.ChangeUserFullNameHandler)
	r.PUT("/user/:id/change-address", controllers.ChangeUserAddressHandler)
	r.PUT("/user/:id/change-user-type", controllers.ChangeUserTypeHandler)
}
