package model

import (
	"time"

	"gorm.io/gorm"
)

type EstimationDoc struct {
	gorm.Model
	DocNum string `gorm:"not null;unique" json:"docnum"`
	CompanyID Company
	Client Client
	DueDate time.Time `gorm:"not null" json:"duedate"`
	Product []Product
	Service []Service
}