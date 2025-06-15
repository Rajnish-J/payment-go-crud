package routes

import (
	"crud/src/controller"

	"github.com/gin-gonic/gin"
)

func SetProductRoutes(router *gin.Engine) {
	router.GET("/products", controller.GetAllProducts)
	router.GET("/products/:id", controller.GetProductByID)
	router.POST("/products", controller.CreateProduct)
	router.PUT("/products/:id", controller.UpdateProduct)
	router.DELETE("/products/:id", controller.DeleteProduct)
}