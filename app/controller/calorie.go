package controller

import (
	"kajiLabTeam/DEBUKATU/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCalorieHandler(c *gin.Context) {
	weightIDPass := c.Param("weightId")
	weightID, err := strconv.ParseInt(weightIDPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid weight ID"})
		return
	}

	calorieService := service.CalorieService{}
	calories, err := calorieService.GetCalories(weightID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get weights"})
		return
	}

	c.JSON(http.StatusOK, calories)
}

func CreateCalorieHandler(c *gin.Context) {
	modelIDPass := c.Param("modelId")
	modelID, err := strconv.ParseInt(modelIDPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid model ID"})
		return
	}
	weightIDPass := c.Param("weightId")
	weightID, err := strconv.ParseInt(weightIDPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid weight ID"})
		return
	}

	mustQuery := c.Query("must")
	mustCalorie, err := strconv.ParseInt(mustQuery, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid must calorie"})
		return
	}
	currentQuery := c.Query("current")
	currentCalorie, err := strconv.ParseInt(currentQuery, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid current calorie"})
		return
	}

	modelService := service.CalorieService{}
	id, err := modelService.CreateCalorie(weightID, modelID, mustCalorie, currentCalorie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"model_id": id})
}
