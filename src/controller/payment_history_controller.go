package controller

import (
	"fmt"
	//"encoding/json"
	"net/http"
	"strconv"

	models "crud/src/model"
	"crud/src/service"

	"github.com/gin-gonic/gin"
)

// CreatePayment godoc
// @Summary      Create a new payment
// @Description  Adds a new payment to the system
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        payment  body      models.Payment  true  "Payment to create"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /api/payments [post]
func CreatePayment(c *gin.Context) {
	var p models.Payment
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreatePayment(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Payment created successfully"})
}

// GetAllPayments godoc
// @Summary      Get all payments
// @Description  Retrieves all payments from the database
// @Tags         payments
// @Produce      json
// @Success      200  {array}  models.Payment
// @Router       /api/payments [get]
func GetAllPayments(c *gin.Context) {
	payments, err := service.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

// GetPaymentByID godoc
// @Summary      Get payment by ID
// @Description  Retrieves a payment by their ID
// @Tags         payments
// @Produce      json
// @Param        id   path      int  true  "Payment ID"
// @Success      200  {object}  models.Payment
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/payments/{id} [get]
func GetPaymentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	payment, err := service.GetPaymentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

// UpdatePayment godoc
// @Summary      Update an existing payment
// @Description  Updates payment data by ID
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "Payment ID"
// @Param        payment  body      models.Payment  true  "Updated payment"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /api/payments/{id} [put]
func UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	var uintID uint
	if _, err := fmt.Sscanf(id, "%d", &uintID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	var req models.Payment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the ID from path param is used
	req.ID = uintID

	if err := service.UpdatePayment(uintID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment updated"})
}

// DeletePayment godoc
// @Summary      Delete a payment
// @Description  Removes a payment from the database by ID
// @Tags         payments
// @Produce      json
// @Param        id   path      int  true  "Payment ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/payments/{id} [delete]
func DeletePayment(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if err := service.DeletePayment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
