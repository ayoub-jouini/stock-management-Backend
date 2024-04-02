package models

import "gorm.io/gorm"

type CompanyAdmin struct {
	gorm.Model
	UserID uint `gorm:"primary Key"`
	CompanyID uint `gorm:"primary Key"`

	User User `gorm:"foreignKey:UserID;references:ID"`
	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
}