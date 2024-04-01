package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	CompanyID uint `gorm:"primary Key"`
	CategoryID uint
	Name string `gorm:"size:255;not null;" json:"name"`
	Description string `json:"description"`
	Quantity int32 `json:"quantity"`
	Price float32 `json:"price"`
	State string `json:"state"`

	Company Company `gorm:"foreignKey:CompanyID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}