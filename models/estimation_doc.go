package model

import "gorm.io/gorm"

type EstimationDoc struct {
	gorm.Model
	DocNum string `gorm:"not null;unique" json:"docnum"`
	CompanyID Company
	Client Client
	DueDate string `gorm:"not null" json:"duedate"`
	Product []Product
	Service []Service
}