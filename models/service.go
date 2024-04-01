package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	CompanyID uint `gorm:"primaryKey"`
	CategoryID uint
	Name string `gorm:"size:255;not null;" json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	State string `json:"state"`
	
	Company Company `gorm:"foreignKey:CompnayID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}