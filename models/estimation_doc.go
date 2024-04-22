package models

import (
	"time"

	"gorm.io/gorm"
)

type EstimationDoc struct {
	gorm.Model
	DocNum string `gorm:"primary Key;not null;unique" json:"docnum"`
	CompanyID uint `gorm:"primary Key"`
	ClientID uint
	DueDate time.Time `gorm:"not null" json:"duedate"`
	Product []*Product `gorm:"foreignKey:EstimationDocID"`
	Service []*Service `gorm:"foreignKey:EstimationDocID"`

	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
	Client Client `gorm:"foreignKey:ClientID;references:ID"`
}

type EstimationDocs []EstimationDoc

func (estimationDoc *EstimationDoc) Save() (error) {
	err := database.Database.Create(&estimationDoc).Error
	if err != nil {
		return err
	}
	return nil
}

func (estimationDoc *EstimationDoc) FindById(id string) (error) {
	err := database.Database.Preload("Company").Preload("Category").Preload("Client").Preload("Product").Preload("Service").Where("ID=?", id).Find(&estimationDoc).Error
	if err != nil {
		return err
	}
	return nil
}

func (estimationDocs *EstimationDocs) FindAll(page string, limit string) (error) {
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&estimationDocs).Error
	if err != nil {
		return err
	}
	return nil
}

func (estimationDoc *EstimationDoc) Update() (error) {
	err := database.Database.Save(&estimationDoc).Error
	if err != nil {
		return err
	}
	return nil
}

func (estimationDoc *EstimationDoc) Delete() (error) {
	err := database.Database.Delete(EstimationDoc{}, "ID = ?", estimationDoc.ID).Error
	if err != nil {
		return err
	}
	return nil
}