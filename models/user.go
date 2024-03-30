package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    FirstName string `gorm:"size:255;not null" json:"firstname"`
	LastName string `gorm:"size:255;not null" json:"lastname"`
	Email string `gorm:"size:255;not null;unique" json:"email"`
	Phone string `gorm:"size:10;not null;unique" json:"phone"`
    Password string `gorm:"size:255;not null;" json:"-"`
	Avatar string `json:"avatar"`
}