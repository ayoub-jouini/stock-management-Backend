package model

import "gorm.io/gorm"

type Tax struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	Type string `gorm:"size:255;not null" json:"type"`
	Value string `json:"value"`
	CompanyID uint 
	
	Company Company `gorm:"foreignKey:CompanyID"`
}