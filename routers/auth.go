package routers

import (	
	"example/web-service-gin/user"

	"github.com/gin-gonic/gin"
)

func setAuthRoute(router *gin.Engine) {
	authController := new(userModule.UserController)
	router.POST("/login", authController.Login)
	router.POST("/signup", authController.Signup)
}
