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

		versionEngine.GET("/model/:userId", controller.GetModelHandler)
		versionEngine.POST("/model/:userId", controller.CreateModelHandler)

		versionEngine.GET("/weight/:userId/:modelId", controller.GetWeightHandler)
		versionEngine.POST("/weight/:userId/:modelId", controller.CreateWeightHandler)

		versionEngine.GET("/calorie/:modelId/:weightId", controller.GetCalorieHandler)
		versionEngine.POST("/calorie/:modelId/:weightId", controller.CreateCalorieHandler)
	}

	r.Run(":8090")

}
