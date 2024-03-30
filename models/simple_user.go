package model

import "gorm.io/gorm"

type SimpleUser struct {
	gorm.Model
	UserID User
	CompanyID Company
}