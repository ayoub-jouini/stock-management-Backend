package models

import (
	"stock_management/database"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	RegNum string `gorm:"size:255;not null;unique" json:"regnum"`
	Name string `gorm:"size:255;not null;" json:"name"`
	Description string `json:"description"`
	Email string `gorm:"size:255;" json:"email"`
	Address string `gorm:"size:255;" json:"address"`
	City string `gorm:"size:255;" json:"city"`
	Country string `gorm:"size:255;" json:"country"`
	Phone string `gorm:"size:255;" json:"phone"`
	Logo string	`gorm:"size:255;" json:"logo"`
	Admin uint
	Employees []*User `gorm:"foreignKey:CompanyID"`

	User *User `gorm:"foreignKey:Admin;references:ID"`
}

func (company *Company) Save() (*Company, error) {
	err := database.Database.Create(&company).Error
	if err != nil {
		return &Company{}, err
	}
	return company, nil
}

func FindCompanyByID(id uint) (Company, error) {
	var company Company
	err := database.Database.Preload("Employees").Preload("Admin").Where("ID=?", id).Find(&company).Error
	if err != nil {
		return company{}, err
	}
	return company, nil
}

func FindAllCompanies(page int, limit int) ([]Company, error) {
	
	var companies []models.Company
	
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := Ctr.DB.Limit(intLimit).Offset(offset).Find(&companies).Error
	if err != nil {
		return companies{}, err
	}
	return companies, nil
}