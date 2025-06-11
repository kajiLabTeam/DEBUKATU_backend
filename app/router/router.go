package router

import (
	"kajiLabTeam/DEBUKATU/controller"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	versionEngine := r.Group("/api")
	{
		versionEngine.GET("/users", controller.GetUserHandler)
		versionEngine.POST("/users", controller.CreateUserHandler)
		versionEngine.PUT("/users/:id", controller.UpdateUserHandler)
		versionEngine.GET("/weight/:userId", controller.GetWeightHandler)
		versionEngine.PUT("/weight/:userId", controller.UpdateWeightHandler)
	}

	r.Run(":8090")

}
