package controller

import (
	"kajiLabTeam/DEBUKATU/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	var userID int64
	userID, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user_id Bad Request"})
		userID = 0
	}
	userService := service.UserService{}
	users, err := userService.GetUsers(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

type UpdateUserRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func UpdateUserHandler(c *gin.Context) {
	var req UpdateUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userService := service.UserService{}
	if err := userService.UpdateUser(req.ID, req.Name); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func CreateUserHandler(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	userService := service.UserService{}
	id, err := userService.CreateUser(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_id": id})
}
