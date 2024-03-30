package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	CompanyID Company
	CategoryID Category
	Name string `gorm:"size:255;not null;" json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	State string `json:"state"`
}