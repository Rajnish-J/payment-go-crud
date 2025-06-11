package main

import (
	"github.com/gin-gonic/gin"
	"crud/src/config"
	"crud/src/routes"
)

func main() {
	config.InitDB()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
