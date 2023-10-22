package routes

import (
	"github.com/airbornharsh/holio-go/controllers"
	"github.com/airbornharsh/holio-go/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignupHandler)
	r.POST("/login", controllers.LoginHanlder)
	r.POST("/logout", controllers.LogoutHandler)
	r.POST("/forgot-password", controllers.ForgotPasswordHandler)
	r.POST("/reset-password", controllers.ResetPasswordHandler)
	r.POST("/change-password", middlewares.TokenVerifyMiddleWare, controllers.ChangePasswordHandler)
	r.POST("/change-email", controllers.ChangeEmailHandler)
	r.POST("/change-phone", controllers.ChangePhoneHandler)
}
