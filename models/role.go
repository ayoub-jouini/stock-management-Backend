package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	Description string `json:"description"`
}