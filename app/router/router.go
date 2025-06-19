package router

import (
	"kajiLabTeam/DEBUKATU/controller"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://debukatu-frontend2-for-deploy.vercel.app")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

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
