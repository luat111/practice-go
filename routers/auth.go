package routers

import (
	"example/web-service-gin/middlewares"
	"example/web-service-gin/user"

	"github.com/gin-gonic/gin"
)

func setAuthRoute(router *gin.Engine) {
	authController := new(userModule.AuthController)
	router.POST("/login", authController.Login)
	router.POST("/signup", authController.Signup)

	authGroup := router.Group("/user")
	authGroup.Use(middlewares.Authentication())
	authGroup.GET("/profile", authController.Profile)
}
