package routes

import (
	"crud/src/controller"

	"github.com/gin-gonic/gin"
)

func SetProductRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/products", controller.GetAllProducts)
		api.GET("/products/:id", controller.GetProductByID)
		api.POST("/products", controller.CreateProduct)
		api.PUT("/products/:id", controller.UpdateProduct)
		api.DELETE("/products/:id", controller.DeleteProduct)
	}
}