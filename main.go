package main

import (
	"fmt"
	"stock_management/migrate"
	"stock_management/routes"

	"github.com/gin-gonic/gin"
)

func init() {
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

	routes.BillRoutes(router)

	routes.CategoryRoutes(router)

	routes.ClientRoutes(router)

	routes.CompanyRoutes(router)

	routes.EstimationDocRoutes(router)

	routes.ProductRoutes(router)

	routes.RoleRoutes(router)

	routes.ServiceRoutes(router)

	routes.SuperUserRoutes(router)
	
	routes.TaxRoutes(router)

	server.Run(":8000")
	fmt.Println("Server running on port 8000")
}