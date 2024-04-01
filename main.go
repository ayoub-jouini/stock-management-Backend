package main

import (
    "diary_api/controllers"
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
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/", controllers.AddCompany)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}