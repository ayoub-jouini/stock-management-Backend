package routes

import (
	"stock_management/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	AuthControllers controllers.AuthControllers
}

func AuthRoutesInit(authController controllers.AuthControllers) AuthRouteController {
	return AuthRouteController{authController}
}

func (ctr AuthRouteController) AuthRoutes(routerGroup *gin.RouterGroup) {
	
	router := routerGroup.Group("auth")

	router.POST("/register", ctr.AuthControllers.Register)
	
	router.POST("/login", ctr.AuthControllers.Login)
}