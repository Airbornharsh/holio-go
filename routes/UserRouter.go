package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUserHandler)
	r.GET("/users/:id", controllers.GetUserHandler)
	r.PUT("/users/:id", controllers.UpdateUserHandler)
	r.DELETE("/users/:id", controllers.DeleteUserHandler)
	r.PUT("/users/:id/change-username", controllers.ChangeUserNameHandler)
	r.PUT("/users/:id/change-fullname", controllers.ChangeUserFullNameHandler)
	r.PUT("/users/:id/change-address", controllers.ChangeUserAddressHandler)
	r.PUT("/users/:id/change-user-type", controllers.ChangeUserTypeHandler)
}
