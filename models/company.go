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
	Employees []*User `gorm:"foreignkey:CompanyID;references:ID"`

	User *User `gorm:"foreignkey:Admin;references:ID" json:"-"`
}

type Companies []Company

func (company *Company) Save() (error) {
	err := database.Database.Create(&company).Error
	if err != nil {
		return err
	}
	return nil
}

func (company *Company) FindById(id string) (error) {
	err := database.Database.Preload("Admin").Preload("Employees").Where("ID=?", id).Find(&company).Error
	if err != nil {
		return err
	}
	return nil
}

func (companies Companies) FindAll(page string, limit string) (error) {
	
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&companies).Error
	if err != nil {
		return err
	}
	return nil
}

func (company *Company) Update() (error) {
	err := database.Database.Save(&company).Error
	if err != nil {
		return err
	}
	return nil
}

func (company Company) Delete() (error) {
	err := database.Database.Delete(Company{}, "ID = ?", company.ID).Error
	if err != nil {
		return err
	}
	return nil
}