package models

import (
	"stock_management/database"
	"strconv"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	CompanyID uint `gorm:"primary Key"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Description string `json:"description"`

	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
}

type Categories []Category

func (category *Category) Save() (error) {
	err := database.Database.Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (category *Category) FindById(id string) (error) {
	err := database.Database.Preload("Company").Where("ID=?", id).Find(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (categories *Categories) FindAll(page string, limit string) (error) {
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&categories).Error
	if err != nil {
		return err
	}
	return nil
}

func (category *Category) Update() (error) {
	err := database.Database.Save(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (category *Category) Delete() (error) {
	err := database.Database.Delete(Category{}, "ID = ?", category.ID).Error
	if err != nil {
		return err
	}
	return nil
}