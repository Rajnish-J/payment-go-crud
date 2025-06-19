package routes

import (
	"crud/src/controller"
	"crud/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")
	{
		// Public routes
		api.POST("/users", controller.CreateUser)      // Sign up
		api.POST("/users/login", controller.LoginUser) // Login
		api.GET("/users", controller.GetAllUsers)

		// Protected routes
		authorized := api.Group("/users")
		authorized.Use(middleware.AuthMiddleware())
		{
			// authorized.GET("/", controller.GetAllUsers)
			authorized.GET("/:id", controller.GetUserByID)
			authorized.PUT("/:id", controller.UpdateUser)
			authorized.DELETE("/:id", controller.DeleteUser)
			authorized.POST("/placeOrder", controller.PlaceOrder)
		}
	}
}

// Create user
//     curl -X POST http://localhost:8080/api/users -H "Content-Type: application/json" -d "{\"id\":3,\"name\":\"Shyam\",\"email\":\"Ram.doe@example.com\",\"password\":\"Password@123\",\"age\":30,\"address\":\"Jaipur, Rajasthan\",\"phone\":\"9876543210\"}"

// Login curl
//      curl -X POST http://localhost:8080/api/users/login -H "Content-Type: application/json" -d "{\"email\":\"Ram.doe@example.com\",\"password\":\"Password@123\"}"

// GetBYID curl
//      curl -X GET http://localhost:8080/api/users/2 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJyb2xlIjoidXNlciIsImV4cCI6MTc1MDQwNDE0MX0._1OsXdAsmu53zdRhnjAPidq24oRowC9uhXFJ_5Ezgu0"

// GetALLUSers
//	   curl -X GET http://localhost:8080/api/users
