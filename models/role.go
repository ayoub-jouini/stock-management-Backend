package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	Description string `json:"description"`
	UserID uint

	User User `gorm:"foreignKey:UserID;references:ID"`
}