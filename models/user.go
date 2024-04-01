package model

import  (
    "diary_api/database"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "html"
    "strings"
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
	Role []Role
}

func (user *User) Save() (*User, error) {
	err :=database.Database.Create(&user).error
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