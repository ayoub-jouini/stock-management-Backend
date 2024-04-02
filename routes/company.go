package routes

import (
	"github.com/gin-gonic/gin"
	"stock_management/controllers"
	"stock_management/middlewares"
)

type CompanyRouteController struct {
	companyController controllers.CompanyController
}

func CompanyRoutesInit(companyController controllers.CompanyController) CompanyRouteController {
	return CompanyRouteController{companyController}
}

func (ctr CompanyRouteController) CompanyRoutes(routerGroup *gin.RouterGroup) {
	
	router := routerGroup.Group("company")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", ctr.companyController.GetAllCompanies)

	router.GET("/:id", ctr.companyController.GetCompanyByID)

	router.POST("/", ctr.companyController.AddCompany)
}