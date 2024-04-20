package routes

import (
	"stock_management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routerGroup *gin.RouterGroup){
	
	router := routerGroup.Group("user")

	router.GET("/", controllers.GetAllUsers)

	router.GET("/:id", controllers.GetUserByID)

	router.PATCH("/:id", controllers.UpdateUserById)
}