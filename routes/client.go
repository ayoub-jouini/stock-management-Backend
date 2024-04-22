package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

func ClientRoutes(routerGroup *gin.RouterGroup) {

	router := routerGroup.Group("client")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllClients)

	router.GET("/:id", controllers.GetClientByID)

	router.POST("/", controllers.AddClient)

	router.PATCH("/:id", controllers.UpdateClient)

	router.DELETE("/:id", controllers.DeleteClient)
}