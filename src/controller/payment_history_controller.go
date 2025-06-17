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

func GetAllPayments(c *gin.Context) {
	payments, err := service.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

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

func DeletePayment(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if err := service.DeletePayment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
