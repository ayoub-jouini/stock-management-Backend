package model

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
	Product []Product
	Service []Service

	Company Company `gorm:"foreignKey:CompanyID"`
	Clinet Client `gorm:"foreignKey:ClientID"`
}