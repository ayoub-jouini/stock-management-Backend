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
	Employee []*User `gorm:"foreignKey:CompanyID"`

	User User `gorm:"foreignKey:Admin;references:ID"`
}

func (company *Company) Save() (*Company, error) {
	err := database.Database.Create(&company).Error
	if err != nil {
		return &Company{}, err
	}
	return company, nil
}