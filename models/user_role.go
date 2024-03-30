package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserID User
	RoleID Role
}