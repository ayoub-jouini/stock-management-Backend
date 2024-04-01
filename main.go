package main

import (
    "diary_api/controller"
    "diary_api/database"
    "diary_api/model"
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
	database.loadEnv()
	database.LoadDatabase()
	serverApplication()
}

func serverApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}