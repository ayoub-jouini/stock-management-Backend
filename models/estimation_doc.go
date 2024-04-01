package model

import (
	"time"

	"gorm.io/gorm"
)

type EstimationDoc struct {
	gorm.Model
	DocNum string `gorm:"primary Key;not null;unique" json:"docnum"`
	CompanyID uint `gorm:"primary Key"`
	ClientID uint
	DueDate time.Time `gorm:"not null" json:"duedate"`
	Product []Product
	Service []Service

	Company Compnay `gorm:"foreignKey:CompanyID"`
	Client Client `gorm:"foreignKey:ClientID"`
}