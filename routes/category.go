package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(routerGroup *gin.RouterGroup) {

	router := routerGroup.Group("category")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllCategories)

	router.GET("/:id", controllers.GetCategoryByID)

	router.POST("/", controllers.AddCategory)

	router.PATCH("/:id", controllers.UpdateCategory)

	router.DELETE("/:id", controllers.DeleteCategory)
}