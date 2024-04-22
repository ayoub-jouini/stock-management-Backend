package models

import "gorm.io/gorm"

type Tax struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	Type string `gorm:"size:255;not null" json:"type"`
	Value string `json:"value"`
	CompanyID uint 
	
	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
}

type Taxs []Tax

func (tax *Tax) Save() (error) {
	err := database.Database.Create(&tax).Error
	if err != nil {
		return err
	}
	return nil
}

func (tax *Tax) FindById(id string) (error) {
	err := database.Database.Preload("Company").Where("ID=?", id).Find(&tax).Error
	if err != nil {
		return err
	}
	return nil
}

func (taxs *Taxs) FindAll(page string, limit string) (error) {
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&taxs).Error
	if err != nil {
		return err
	}
	return nil
}

func (tax *Tax) Update() (error) {
	err := database.Database.Save(&tax).Error
	if err != nil {
		return err
	}
	return nil
}

func (tax *Tax) Delete() (error) {
	err := database.Database.Delete(Tax{}, "ID = ?", tax.ID).Error
	if err != nil {
		return err
	}
	return nil
}