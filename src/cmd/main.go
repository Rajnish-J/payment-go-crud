// @title           Payment CRUD API
// @version         1.0
// @description     Swagger docs for user, product, company, payment history APIs
// @host            localhost:8080
// @BasePath        /

package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "crud/docs"

	"crud/src/config"
	"crud/src/routes"
)

func main() {
	config.InitDB()
	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	routes.SetupRoutes(router)
	routes.SetProductRoutes(router)
	routes.RegisterPaymentRoutes(router)
	routes.RegisterCompanyRoutes(router)

	router.Run(":8080")
}
