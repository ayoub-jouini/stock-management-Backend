package routes

import (
	"github.com/gin-gonic/gin"
	"stock_management/controllers"
)

type AuthRouteController struct {
	AuthControllers controllers.AuthControllers
}

func AuthRoutesInit(authController controllers.AuthControllers) AuthRouteController {
	return AuthRouteController{authController}
}

func AuthRoutes(routerGroup *gin.RouterGroup) {
	
	router := routerGroup.Group("auth")

	router.POST("/register", controllers.Register)
	
	router.POST("/login", controllers.Login)
}