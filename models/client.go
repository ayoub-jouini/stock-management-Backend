package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	Avatar string `json:"avatar"`
}