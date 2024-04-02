package models

import (
	"stock_management/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
    FirstName string `gorm:"size:255;not null" json:"firstname"`
	LastName string `gorm:"size:255;not null" json:"lastname"`
	Email string `gorm:"size:255;not null;unique" json:"email"`
	Phone string `gorm:"size:10;not null;unique" json:"phone"`
    Password string `gorm:"size:255;not null;" json:"-"`
	Avatar string `json:"avatar"`
	Role []Role `gorm:"foreignKey:UserID"`
	
	CompanyID uint 

	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
}

func (user *User) Save() (*User, error) {
	err :=database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSava(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByEmail(email string) (User, error) {
	var user User
	err := database.Database.Where("email=?", email).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Role").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}