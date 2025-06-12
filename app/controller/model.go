package controller

import (
	"kajiLabTeam/DEBUKATU/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetModelHandler(c *gin.Context) {
	userIDPass := c.Param("userId")
	userID, err := strconv.ParseInt(userIDPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	modelService := service.ModelService{}
	models, err := modelService.GetModels(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
		return
	}

	c.JSON(http.StatusOK, models)
}

func CreateModelHandler(c *gin.Context) {
	userIDPass := c.Param("userId")
	userID, err := strconv.ParseInt(userIDPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	weightQuery := c.Query("weight")
	weight, err := strconv.ParseFloat(weightQuery, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid weight"})
		return
	}
	monthQuery := c.Query("month")
	month, err := strconv.ParseFloat(monthQuery, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid month"})
		return
	}

	modelService := service.ModelService{}
	id, err := modelService.CreateModel(userID, weight, month)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"model_id": id})
}
