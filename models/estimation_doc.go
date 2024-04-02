package models

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
	Product []*Product `gorm:"foreignKey:EstimationDocID"`
	Service []*Service `gorm:"foreignKey:EstimationDocID"`

	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
	Client Client `gorm:"foreignKey:ClientID;references:ID"`
}