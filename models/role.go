package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
	Description string `json:"description"`
}