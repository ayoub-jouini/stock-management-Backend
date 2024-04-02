package models

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	BillNum string `gorm:"primary Key;not null;unique" json:"billnum"`
	CompanyID uint `gorm:"primary Key"`
	ClientID uint
	DueDate time.Time `gorm:"not null" json:"duedate"`
	Product []*Product `gorm:"foreignKey:BillID"`
	Service []*Service `gorm:"foreignKey:BillID"`

	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
	Clinet Client `gorm:"foreignKey:ClientID;references:ID"`
}