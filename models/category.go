package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CompanyID Company
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Description string `json:"description"`
}