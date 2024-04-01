package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	CompanyID uint `gorm:"primary Key"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Description string `json:"description"`

	Company Company `gorm:"foreignKey:CompanyID"`
}