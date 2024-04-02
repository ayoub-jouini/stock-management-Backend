package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

type CompanyRouteController struct {
	companyController controllers.CompanyControllers
}

func CompanyRoutesInit(companyController controllers.CompanyControllers) CompanyRouteController {
	return CompanyRouteController{companyController}
}

func (ctr CompanyRouteController) CompanyRoutes(routerGroup *gin.RouterGroup) {
	
	router := routerGroup.Group("company")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", ctr.companyController.GetAllCompanies)

	router.GET("/:id", ctr.companyController.GetCompanyByID)

	router.POST("/", ctr.companyController.AddCompany)
}