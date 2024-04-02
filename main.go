package main

import (
	"fmt"
	"log"
	"stock_management/controllers"
	"stock_management/database"
	"stock_management/middlewares"
	"stock_management/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()
	LoadDatabase()
	serverApplication()
}

func LoadEnv() {
	err  := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}

func LoadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Company{})
	database.Database.AutoMigrate(&models.Category{})
	database.Database.AutoMigrate(&models.Product{})
	database.Database.AutoMigrate(&models.Service{})
	database.Database.AutoMigrate(&models.Client{})
	database.Database.AutoMigrate(&models.Role{})
	database.Database.AutoMigrate(&models.SimpleUser{})
	database.Database.AutoMigrate(&models.CompanyAdmin{})
	database.Database.AutoMigrate(&models.SuperUser{})
	database.Database.AutoMigrate(&models.EstimationDoc{})
	database.Database.AutoMigrate(&models.Bill{})
	database.Database.AutoMigrate(&models.Tax{})
}

func serverApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middlewares.JWTAuthMiddleware())
	protectedRoutes.POST("/", controllers.AddCompany)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}