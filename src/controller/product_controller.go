package controller

import (
	"crud/src/dto"
	"crud/src/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)



// GetAllProducts godoc
// @Summary      Get all products
// @Description  Retrieves all products from the database
// @Tags         products
// @Produce      json
// @Success      200  {array}  models.Product
// @Router       /api/products [get]
func GetAllProducts(c *gin.Context) {
	products, err := service.GetAllProduts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}



// GetProductByID godoc
// @Summary      Get product by ID
// @Description  Retrieves a product by their ID
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/productss/{id} [get]
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



// CreateProduct godoc
// @Summary      Create a new product
// @Description  Adds a new product to the system
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        user  body      models.Product  true  "Product to create"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /api/products [post]
func CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.CreateProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product Created"})
}



// UpdateProduct godoc
// @Summary      Update an existing product
// @Description  Updates product data by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "Product ID"
// @Param        product  body      models.Product  true  "Updated product"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /api/products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var uintID uint
	if _, err := fmt.Sscanf(id, "%d", &uintID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.UpdateProduct(uintID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})

}




// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Removes a product from the database by ID
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var uintID uint
	if _, err := fmt.Sscanf(id, "%d", &uintID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	err := service.DeleteProduct(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
