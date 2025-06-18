package controller

import (
	"kajiLabTeam/DEBUKATU/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	idQuery := c.Query("id")
	nameQuery := c.Query("name")
	passQuery := c.Query("password")

	userService := service.UserService{}

	if idQuery != "" {
		parsedID, err := strconv.ParseInt(idQuery, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
			return
		}

		users, err := userService.GetUsers(parsedID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
			return
		}
		c.JSON(http.StatusOK, users)
		return
	}

	// クエリが全て空 → GetUsers(0)
	if nameQuery == "" && passQuery == "" {
		users, err := userService.GetUsers(0)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
			return
		}
		c.JSON(http.StatusOK, users)
		return
	}

	// name と password の両方がある場合 → GetUser
	if nameQuery != "" && passQuery != "" {
		user, err := userService.GetUser(nameQuery, passQuery)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid name or password"})
			return
		}
		c.JSON(http.StatusOK, user)
		return
	}

	// name か password のどちらか一方だけ → エラー
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "both name and password are required"})
}

type UpdateUserRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func UpdateUserHandler(c *gin.Context) {
	userIdPass := c.Param("id")
	userId, err := strconv.ParseInt(userIdPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	// クエリパラメータ取得
	name := c.Query("name")
	deletedStr := c.Query("deleted")

	// どちらも指定されていない場合はエラー
	if name == "" && deletedStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no update parameter provided"})
		return
	}

	var deleted *bool
	if deletedStr != "" {
		val, err := strconv.ParseBool(deletedStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid deleted value"})
			return
		}
		deleted = &val
	}

	userService := service.UserService{}
	if err := userService.UpdateUser(userId, name, deleted); err != nil {
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
	password := c.Query("password")
	ageQuery := c.Query("age")
	gender := c.Query("gender")
	if gender != "woman" && gender != "man" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "gender is required"})
		return
	}
	age, err := strconv.ParseInt(ageQuery, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "age is required"})
		return
	}

	heightQuery := c.Query("height")
	height, err := strconv.ParseFloat(heightQuery, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "height is required"})
		return
	}

	userService := service.UserService{}
	id, err := userService.CreateUser(name, password, age, height, gender)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_id": id})
}
