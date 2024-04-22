package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

func TaxRoutes(routerGroup *gin.RouterGroup) {

	router := routerGroup.Group("tax")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllTaxes)

	router.GET("/:id", controllers.GetTaxByID)

	router.POST("/", controllers.AddTax)

	router.PATCH("/:id", controllers.UpdateTax)

	router.DELETE("/:id", controllers.DeleteTax)
}