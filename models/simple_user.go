package models

import "gorm.io/gorm"

type SimpleUser struct {
	gorm.Model
	UserID uint `gorm:"primaryKey"`
	CompanyID uint `gorm:"primaryKey"`
	
	User User `gorm:"foreignKey:UserID;references:ID"`
	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
}