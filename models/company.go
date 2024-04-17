package models

import (
	"stock_management/database"
	"strconv"

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
	Phone string `gorm:"size:8;" json:"phone"`
	Logo string	`json:"logo"`
	Admin uint `json:"admin"`
	Employees []*User `gorm:"foreignkey:CompanyID"`

	User *User `gorm:"foreignkey:Admin;references:ID" json:"-"`
}

func (company *Company) Save() (*Company, error) {
	err := database.Database.Create(&company).Error
	if err != nil {
		return &Company{}, err
	}
	return company, nil
}

func FindCompanyByID(id string) (Company, error) {
	var company Company
	err := database.Database.Preload("Employees").Preload("Admin").Where("ID=?", id).Find(&company).Error
	if err != nil {
		return company, err
	}
	return company, nil
}

func FindAllCompanies(page *string, limit *string) ([]Company, error) {
	
	var companies []Company
	
	intPage, _ := strconv.Atoi(*page)
	intLimit, _ := strconv.Atoi(*limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&companies).Error
	if err != nil {
		return companies, err
	}
	return companies, nil
}

func (company *Company) UpdateCompany() (error) {
	err := database.Database.Save(&company).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCompany(id string) (error) {
	err := database.Database.Delete(Company{}, "ID = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}