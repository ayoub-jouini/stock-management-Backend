package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	CompanyID uint `gorm:"primary Key"`
	Name string `gorm:"size:255;not null" json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	Avatar string `json:"avatar"`

	Compnay Compnay `gorm:"foreignKey:CompnayID"`
}