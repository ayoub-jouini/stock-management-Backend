package routes

import (
	"stock_management/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(routerGroup *gin.RouterGroup) {
	
	router := routerGroup.Group("auth")

	router.POST("/register", controllers.Register)
	
	router.POST("/login", controllers.Login)
}