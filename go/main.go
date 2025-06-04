package main

import (
	controller "debukatu_backend/app/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	versionEngine := r.Group("/api")
	{
		versionEngine.GET("/users", controller.GetUserHandler)
		versionEngine.GET("/weight/:userId", controller.GetWeightHandler)
		versionEngine.PUT("/weight/:userId", controller.UpdateWeightHandler)
	}

	r.Run(":8080")
}
