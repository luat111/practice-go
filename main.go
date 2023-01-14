package main

import (
	"example/web-service-gin/routers"
	"example/web-service-gin/utils"
)

func main() {
	router := routers.InitRoute()
	port := utils.EnvVar("SERVER_PORT", "8080")
	router.Run("localhost:" + port)
}
