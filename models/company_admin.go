package model

import "gorm.io/gorm"

type CompanyAdmin struct {
	gorm.Model
	UserID User
	CompanyID Company
}