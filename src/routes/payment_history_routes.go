package routes

import (
	"crud/src/controller"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/payments", controller.CreatePayment)
		api.GET("/payments", controller.GetAllPayments)
		api.GET("/payments/:id", controller.GetPaymentByID)
		api.PUT("/payments/:id", controller.UpdatePayment)
		api.DELETE("/payments/:id", controller.DeletePayment)
	}
}
