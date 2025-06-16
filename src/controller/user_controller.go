package controller

import (
	"net/http"

	models "crud/src/model"
	"crud/src/service"
	"fmt"

	"github.com/gin-gonic/gin"
) 

func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var uintID uint
	if _, err := fmt.Sscanf(id, "%d", &uintID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := service.GetUserByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var uintID uint
	if _, err := fmt.Sscanf(id, "%d", &uintID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.UpdateUser(uintID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	// Convert id from string to uint
	var uintID uint
	if _, err := fmt.Sscanf(id, "%d", &uintID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err := service.DeleteUser(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
