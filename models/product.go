package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CompanyID Company
	CategoryID Category
	Name string `gorm:"size:255;not null;" json:"name"`
	Description string `json:"description"`
	Quantity int32 `json:"quantity"`
	Price float32 `json:"price"`
	State string `json:"state"`
}