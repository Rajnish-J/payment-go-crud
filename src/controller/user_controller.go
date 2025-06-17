package controller

import (
	"net/http"

	models "crud/src/model"
	"crud/src/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieves all users from the database
// @Tags         users
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /api/users [get]
func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Retrieves a user by their ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/users/{id} [get]
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

// CreateUser godoc
// @Summary      Create a new user
// @Description  Adds a new user to the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User to create"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /api/users [post]
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

// UpdateUser godoc
// @Summary      Update an existing user
// @Description  Updates user data by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "User ID"
// @Param        user  body      models.User  true  "Updated user"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /api/users/{id} [put]
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

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Removes a user from the database by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/users/{id} [delete]
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

// PlaceOrder godoc
// @Summary      Place an order
// @Description  User places an order for a product with optional rewards
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        order  body  models.Payment  true  "Order request"
// @Success      201    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Router       /api/users/placeOrder [post]
func PlaceOrder(c *gin.Context) {
	var req models.Payment

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid input: " + err.Error(),
		})
		return
	}

	err := service.PlaceOrderService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Order placed successfully",
	})
}
