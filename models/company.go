package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	RegNum string `gorm:"size:255;not null;unique" json:"regnum"`
	Name string `gorm:"size:255;not null;" json:"name"`
	Description string `json:"description"`
	Email string `gorm:"size:255;" json:"email"`
	Address string `gorm:"size:255;" json:"address"`
	City string `gorm:"size:255;" json:"city"`
	Country string `gorm:"size:255;" json:"country"`
	Phone string `gorm:"size:255;" json:"phone"`
	Logo string	`gorm:"size:255;" json:"logo"`
}