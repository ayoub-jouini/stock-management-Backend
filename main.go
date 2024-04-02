package main

import (
	"fmt"
	"stock_management/routes"
	"stock_management/migrate"
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

	CompanyController = controllers.CompanyControllersInit()
	CompanyRouteController = routes.CompanyRoutesInit(CompanyController)

	AuthController = controllers.CompanyControllersInit()
	AuthRouteController = routes.CompanyRoutesInit(AuthController)
}

func main() {
	migrate.LoadEnv()
	migrate.LoadDatabase()

	serverApplication()
}

func serverApplication() {
	router := gin.Default()

	AuthRouteController.AuthRoutes(router)

	protectedRoutes := router.Group("/api")

	CompanyRouteController.CompanyRoutes(protectedRoutes)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}