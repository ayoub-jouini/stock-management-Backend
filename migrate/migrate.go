package migrate

import (
	"log"
	"stock_management/database"
	"stock_management/models"

	"github.com/joho/godotenv"
)

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
	database.Database.AutoMigrate(&models.SuperUser{})
	database.Database.AutoMigrate(&models.EstimationDoc{})
	database.Database.AutoMigrate(&models.Bill{})
	database.Database.AutoMigrate(&models.Tax{})
}