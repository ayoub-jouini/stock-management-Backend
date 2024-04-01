package model

import "gorm.io/gorm"

type CompanyAdmin struct {
	gorm.Model
	UserID uint `gorm:"primary Key"`
	CompanyID uint `gorm:"primary Key"`

	User User `gorm:"foreignKey:UserID"`
	Company Company `gorm:"foreignKey:CompanyID"`
}