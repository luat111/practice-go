package routers

import (
	"example/web-service-gin/middlewares"
	"example/web-service-gin/user"

	"github.com/gin-gonic/gin"
)

func setUserRoute(router *gin.Engine) {
	userController := new(userModule.UserController)

	userGroup := router.Group("/user")
	userGroup.Use(middlewares.Authentication())
	userGroup.GET("/profile", userController.Profile)
	userGroup.PUT("/:id", userController.Update)
	userGroup.DELETE("/:id", userController.Remove)
}
