package main

import (
	"example/web-service-gin/albums"
	"github.com/gin-gonic/gin"
)




func main() {
	// gin.SetMode(gin.ReleaseMode)
	// router := gin.New()

	router := gin.Default()

	router.SetTrustedProxies([]string{"192.168.0.38"})

	v1 := router.Group("/api")

	albums.AlbumRegister(v1.Group("/albums"))

	router.Run("localhost:8080")
}