package main

import (
	"fmt"
	"stock_management/routes"
	"stock_management/migrate"
	"github.com/gin-gonic/gin"
)

func main() {
	migrate.LoadEnv()
	migrate.LoadDatabase()
	serverApplication()
}

func serverApplication() {
	router := gin.Default()

	routes.AuthRoutes(router)

	protectedRoutes := router.Group("/api")

	routes.CompanyRoutes(protectedRoutes)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}