package routes

import (
	"github.com/gin-gonic/gin"
	"stock_management/controllers"
	"stock_management/middlewares"
)

func CompanyRoutes(routerGroup *gin.RouterGroup) {
	
	router := routerGroup.Group("company")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllCompanies)

	router.GET("/:id", controllers.GetCompanyByID)

	router.POST("/", controllers.AddCompany)
}