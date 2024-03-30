package model

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	BillNum string `gorm:"not null;unique" json:"billnum"`
	CompanyID Company
	Client Client
	DueDate string `gorm:"not null" json:"duedate"`
	Product []Product
	Service []Service
}