package main

import (
	"log"
	"stock_management/database"
	model "stock_management/models"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
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