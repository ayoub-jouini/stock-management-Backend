package migrate

import (
	"log"
	"github.com/joho/godotenv"
	"stock_management/database"
	"stock_management/models"
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
	database.Database.AutoMigrate(&models.SimpleUser{})
	database.Database.AutoMigrate(&models.CompanyAdmin{})
	database.Database.AutoMigrate(&models.SuperUser{})
	database.Database.AutoMigrate(&models.EstimationDoc{})
	database.Database.AutoMigrate(&models.Bill{})
	database.Database.AutoMigrate(&models.Tax{})
}