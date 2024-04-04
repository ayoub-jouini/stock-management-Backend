package main

import (
	"fmt"
	"stock_management/migrate"
	"stock_management/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	//check ken yetfas5ou wala yo9o3dou 
	migrate.LoadEnv()
	migrate.LoadDatabase()
}

func main() {
	serverApplication()
}

func serverApplication() {
	server := gin.Default()

	router := server.Group("/api")

	routes.AuthRoutes(router)

	routes.UserRoutes(router)

	routes.CompanyRoutes(router)

	server.Run(":8000")
	fmt.Println("Server running on port 8000")
}