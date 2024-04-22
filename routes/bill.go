package routes

import (
	"stock_management/controllers"
	"stock_management/middlewares"

	"github.com/gin-gonic/gin"
)

func BillRoutes(routerGroup *gin.RouterGroup) {

	router := routerGroup.Group("bill")

	router.Use(middlewares.JWTAuthMiddleware())

	router.GET("/", controllers.GetAllBills)

	router.GET("/:id", controllers.GetBillByID)

	router.POST("/", controllers.AddBill)

	router.PATCH("/:id", controllers.UpdateBill)

	router.DELETE("/:id", controllers.DeleteBill)
}