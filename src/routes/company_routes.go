package routes

import (
	"crud/src/controller"
	"github.com/gin-gonic/gin"
)

func RegisterCompanyRoutes(router *gin.Engine) {
	company := router.Group("/api/company")
	{
		company.POST("/", controller.CreateCompany)
		company.GET("/", controller.GetCompanies)
		company.GET("/:id", controller.GetCompanyByID)
		company.PUT("/", controller.UpdateCompany)
		company.DELETE("/:id", controller.DeleteCompany)
	}
}
