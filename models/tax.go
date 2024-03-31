package model

import "gorm.io/gorm"

type Tax struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
	Type string `gorm:"size:255;not null" json:"type"`
	Value string `json:"value"`
	CompanyID uint 
	Company Company `gorm:"foreignKey:CompanyID"`
}