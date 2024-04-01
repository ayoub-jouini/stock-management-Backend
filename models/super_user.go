package model

import "gorm.io/gorm"

type SuperUser struct {
	gorm.Model
	UserID uint `gorm:"primaryKey"`
	
	User User `gorm:"foreignKey:UserID"`
}