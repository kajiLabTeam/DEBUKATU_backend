package controller

import (
	"kajiLabTeam/DEBUKATU/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetWeightHandler(c *gin.Context) {
	userIDPass := c.Param("userId")
	userID, err := strconv.ParseInt(userIDPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	modelIDPass := c.Param("modelId")
	modelID, err := strconv.ParseInt(modelIDPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid model ID"})
		return
	}
	weightService := service.WeightService{}
	weights, err := weightService.GetWeights(userID, modelID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get weights"})
		return
	}

	c.JSON(http.StatusOK, weights)
}
func UpdateWeightHandler(c *gin.Context) {
}
