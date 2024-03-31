package model

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	BillNum string `gorm:"not null;unique" json:"billnum"`
	CompanyID Company
	Client Client
	DueDate time.Time `gorm:"not null" json:"duedate"`
	Product []Product
	Service []Service
}