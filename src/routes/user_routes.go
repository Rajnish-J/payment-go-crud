package routes

import (
	"crud/src/controller"
	
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/users", controller.GetAllUsers)
		api.GET("/users/:id", controller.GetUserByID)
		api.POST("/users", controller.CreateUser)
		api.PUT("/users/:id", controller.UpdateUser)
		api.DELETE("/users/:id", controller.DeleteUser)
	}
}

