package routes

import (
	"github.com/gin-gonic/gin"
	"crud/src/controller"
)

func RegisterPaymentRoutes(r *gin.Engine) {
	payment := r.Group("/api/payments")
	{
		payment.POST("/", controller.CreatePayment)
		payment.GET("/", controller.GetPayments)
		payment.GET("/:id", controller.GetPaymentByID)
		payment.PUT("/", controller.UpdatePayment)
		payment.DELETE("/:id", controller.DeletePayment)
	}
}
