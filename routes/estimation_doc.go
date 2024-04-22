package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

func ClientRoutes(routerGroup *gin.RouterGroup) {

	router := routerGroup.Group("estimationdoc")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllEstimationDocs)

	router.GET("/:id", controllers.GetEstimationDocByID)

	router.POST("/", controllers.AddEstimationDoc)

	router.PATCH("/:id", controllers.UpdateEstimationDoc)

	router.DELETE("/:id", controllers.DeleteEstimationDoc)
}