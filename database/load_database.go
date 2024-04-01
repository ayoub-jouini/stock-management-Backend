package database

func LoadDatabase() {
	Connect()
	Database.AutoMigrate(&model.User{})
	Database.AutoMigrate(&model.Company{})
	Database.AutoMigrate(&model.Category{})
	Database.AutoMigrate(&model.Product{})
	Database.AutoMigrate(&model.Service{})
	Database.AutoMigrate(&model.Client{})
	Database.AutoMigrate(&model.Role{})
	Database.AutoMigrate(&model.SimpleUser{})
	Database.AutoMigrate(&model.CompanyAdmin{})
	Database.AutoMigrate(&model.SuperUser{})
	Database.AutoMigrate(&model.EstimationDoc{})
	Database.AutoMigrate(&model.Bill{})
	Database.AutoMigrate(&model.Tax{})
}