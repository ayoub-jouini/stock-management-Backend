package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(routerGroup *gin.RouterGroup) {

	router := routerGroup.Group("role")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllRoles)

	router.GET("/:id", controllers.GetRoleByID)

	router.POST("/", controllers.AddRoles)

	router.PATCH("/:id", controllers.UpdateRole)

	router.DELETE("/:id", controllers.DeleteRole)
}