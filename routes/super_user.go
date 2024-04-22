package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

func ServiceRoutes(routerGroup *gin.RouterGroup) {

	router := routerGroup.Group("superuser")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllSuperUsers)

	router.GET("/:id", controllers.GetSuperUserByID)

	router.POST("/", controllers.AddSuperUser)

	router.PATCH("/:id", controllers.UpdateSuperUser)

	router.DELETE("/:id", controllers.DeleteSuperUser)
}