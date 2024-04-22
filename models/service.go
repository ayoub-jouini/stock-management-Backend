package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	CompanyID uint `gorm:"primaryKey"`
	CategoryID uint
	Name string `gorm:"size:255;not null;" json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	State string `json:"state"`

	BillID uint
	EstimationDocID uint

	Bill Bill `gorm:"foreignKey:BillID;references:ID"`
	EstimationDoc EstimationDoc `gorm:"foreignKey:EstimationDocID;references:ID"`
	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
}

type Services []Service 

func (service Service) Save() (error) {
	err := database.Database.Create(&service).Error
	if err != nil {
		return err
	}
	return nil
}

func (service Service) FindById(id string) (error) {
	err := database.Database.Preload("Company").Preload("Category").Where("ID=?", id).Find(&service).Error
	if err != nil {
		return err
	}
	return nil
}

func (services Services) FindAll(page string, limit string) (error) {
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&services).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) Update() (error) {
	err := database.Database.Save(&service).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) Delete() (error) {
	err := database.Database.Delete(Service{}, "ID = ?", service.ID).Error
	if err != nil {
		return err
	}
	return nil
}