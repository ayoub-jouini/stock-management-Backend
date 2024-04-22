package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

func ServiceRoutes(routerGroup *gin.RouterGroup) {

	router := routerGroup.Group("service")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllServices)

	router.GET("/:id", controllers.GetServiceByID)

	router.POST("/", controllers.AddService)

	router.PATCH("/:id", controllers.UpdateService)

	router.DELETE("/:id", controllers.DeleteService)
}