package controller

import (
	"crud/src/dto"
	"crud/src/service"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var productId uint
	if _, err := fmt.Sscanf(id, "%d", &productId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product Id"})
		return
	}
	product, err := service.GetProductByID(productId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err:=c.ShouldBindJSON(&req); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err:=service.CreateProduct(req)
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"Product Created"})
}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
