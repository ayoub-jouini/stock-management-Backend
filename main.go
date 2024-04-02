package main

import (
	"fmt"
	"stock_management/controllers"
	"stock_management/database"
	"stock_management/migrate"
	"stock_management/routes"

	"github.com/gin-gonic/gin"
)

var (
	AuthController controllers.AuthControllers
	AuthRouteController routes.AuthRouteController

	CompanyController      controllers.CompanyControllers
	CompanyRouteController routes.CompanyRouteController
)

func init() {
	//check ken yetfas5ou wala yo9o3dou 
	migrate.LoadEnv()
	migrate.LoadDatabase()

	AuthController = controllers.AuthControllersInit(database.Database)
	AuthRouteController = routes.AuthRoutesInit(AuthController)

	CompanyController = controllers.CompanyControllersInit(database.Database)
	CompanyRouteController = routes.CompanyRoutesInit(CompanyController)
}

func main() {
	serverApplication()
}

func serverApplication() {
	server := gin.Default()

	router := server.Group("/api")

	AuthRouteController.AuthRoutes(router)

	CompanyRouteController.CompanyRoutes(router)

	server.Run(":8000")
	fmt.Println("Server running on port 8000")
}