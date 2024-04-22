package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	CompanyID uint `gorm:"primary Key"`
	CategoryID uint
	Name string `gorm:"size:255;not null;" json:"name"`
	Description string `json:"description"`
	Quantity int32 `json:"quantity"`
	Price float32 `json:"price"`
	State string `json:"state"`

	BillID uint
	EstimationDocID uint

	Bill Bill `gorm:"foreignKey:BillID;references:ID"`
	EstimationDoc EstimationDoc `gorm:"foreignKey:EstimationDocID;references:ID"`
	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
}

type Products []Product 

func (product Product) Save() (error) {
	err := database.Database.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (product Product) FindById(id string) (error) {
	err := database.Database.Preload("Company").Preload("Category").Where("ID=?", id).Find(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (products Products) FindAll(page string, limit string) (error) {
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&products).Error
	if err != nil {
		return err
	}
	return nil
}

func (product *Product) Update() (error) {
	err := database.Database.Save(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (product *Product) Delete() (error) {
	err := database.Database.Delete(Product{}, "ID = ?", product.ID).Error
	if err != nil {
		return err
	}
	return nil
}