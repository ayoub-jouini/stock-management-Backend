package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	Description string `json:"description"`
	UserID uint

	User User `gorm:"foreignKey:UserID;references:ID"`
}

type Roles []Role 

func (role Role) Save() (error) {
	err := database.Database.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (role Role) FindById(id string) (error) {
	err := database.Database.Where("ID=?", id).Find(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (roles Roles) FindAll(page string, limit string) (error) {
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&roles).Error
	if err != nil {
		return err
	}
	return nil
}

func (role *Role) Update() (error) {
	err := database.Database.Save(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (role *Role) Delete() (error) {
	err := database.Database.Delete(Role{}, "ID = ?", role.ID).Error
	if err != nil {
		return err
	}
	return nil
}