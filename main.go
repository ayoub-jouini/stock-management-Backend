package main

import (
    "diary_api/controller"
    "diary_api/database"
    "diary_api/model"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
)

func main() {
	loadEnv()
	loadDatabase()
	serverApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Company{})
	database.Database.AutoMigrate(&model.Category{})
	database.Database.AutoMigrate(&model.Product{})
	database.Database.AutoMigrate(&model.Service{})
	database.Database.AutoMigrate(&model.Client{})
	database.Database.AutoMigrate(&model.Role{})
	database.Database.AutoMigrate(&model.SimpleUser{})
	database.Database.AutoMigrate(&model.CompanyAdmin{})
	database.Database.AutoMigrate(&model.SuperUser{})
	database.Database.AutoMigrate(&model.EstimationDoc{})
	database.Database.AutoMigrate(&model.Bill{})
	database.Database.AutoMigrate(&model.Tax{})
}

func loadEnv() {
	err  := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}

func serverApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}