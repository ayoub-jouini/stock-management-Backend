package models

import (
	"stock_management/database"
	"strconv"

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
    Password string `gorm:"not null;" json:"password"`
	Avatar string `json:"avatar"`
	Role []Role `gorm:"foreignKey:UserID"`
	
	CompanyID uint `gorm:"-" json:"-"`

	// Company Company `gorm:"foreignKey:CompanyID;references:ID"`
}

type Users []User

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSava() error {
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

func (user *User) FindByEmail(email string) (error) {
	err := database.Database.Where("email=?", email).Find(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *User) FindById(id string) (error) {
	err := database.Database.Preload("Role").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (users Users) FindAll(page string, limit string) (error) {
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&users).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Update() (error) {
	err := database.Database.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Delete() (error) {
	err := database.Database.Delete(User{}, "ID = ?", user.ID).Error
	if err != nil {
		return err
	}
	return nil
}