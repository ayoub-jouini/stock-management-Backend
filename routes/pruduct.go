package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(routerGroup *gin.RouterGroup) {

	router := routerGroup.Group("product")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllProducts)

	router.GET("/:id", controllers.GetProductByID)

	router.POST("/", controllers.AddProduct)

	router.PATCH("/:id", controllers.UpdateProduct)

	router.DELETE("/:id", controllers.DeleteProduct)
}