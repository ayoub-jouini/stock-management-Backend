package routes

import (
	"github.com/gin-gonic/gin"
	"stock_management/controllers"
)

func AuthRoutes(routerGroup *gin.RouterGroup) {
	
	router := routerGroup.Group("auth")

	router.POST("/register", controllers.Register)
	
	router.POST("/login", controllers.Login)
}