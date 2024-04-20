package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(routerGroup *gin.RouterGroup) {
	
	router := routerGroup.Group("company")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllCompanies)

	router.GET("/:id", controllers.GetCompanyByID)

	router.POST("/", controllers.AddCompany)

	router.PATCH("/:id", controllers.UpdateCompanyById)
}