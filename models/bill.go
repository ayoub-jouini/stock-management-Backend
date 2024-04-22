package models

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	BillNum string `gorm:"primary Key;not null;unique" json:"billnum"`
	CompanyID uint `gorm:"primary Key"`
	ClientID uint
	DueDate time.Time `gorm:"not null" json:"duedate"`
	Product []*Product `gorm:"foreignKey:BillID"`
	Service []*Service `gorm:"foreignKey:BillID"`

	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
	Client Client `gorm:"foreignKey:ClientID;references:ID"`
}

type Bills []Bill

func (bill *Bill) Save() (error) {
	err := database.Database.Create(&bill).Error
	if err != nil {
		return err
	}
	return nil
}

func (bill *Bill) FindById(id string) (error) {
	err := database.Database.Preload("Company").Preload("Category").Preload("Client").Preload("Product").Preload("Service").Where("ID=?", id).Find(&bill).Error
	if err != nil {
		return err
	}
	return nil
}

func (bills *Bills) FindAll(page string, limit string) (error) {
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&bills).Error
	if err != nil {
		return err
	}
	return nil
}

func (bill *Bill) Update() (error) {
	err := database.Database.Save(&bill).Error
	if err != nil {
		return err
	}
	return nil
}

func (bill *Bill) Delete() (error) {
	err := database.Database.Delete(Bill{}, "ID = ?", bill.ID).Error
	if err != nil {
		return err
	}
	return nil
}