package routes

import (
	"crud/src/controller"
	"github.com/gin-gonic/gin"
)

func RegisterCompanyRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/company", controller.CreateCompany)
		api.GET("/company", controller.GetCompanies)
		api.GET("/company:id", controller.GetCompanyByID)
		api.PUT("/company", controller.UpdateCompany)
		api.DELETE("/company:id", controller.DeleteCompany)
	}
}
