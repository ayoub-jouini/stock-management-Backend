package models

import "gorm.io/gorm"

type SuperUser struct {
	gorm.Model
	UserID uint `gorm:"primaryKey"`
	
	User User `gorm:"foreignKey:UserID;references:ID"`
}

type SuperUsers []SuperUser

func (superUser *SuperUser) Save() (error) {
	err := database.Database.Create(&superUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (superUser *SuperUser) FindById(id string) (error) {
	err := database.Database.Preload("User").Where("ID=?", id).Find(&superUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (superUsers *SuperUsers) FindAll(page string, limit string) (error) {
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Preload("User").Find(&superUsers).Error
	if err != nil {
		return err
	}
	return nil
}

func (superUser *SuperUser) Update() (error) {
	err := database.Database.Save(&superUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (superUser *SuperUser) Delete() (error) {
	err := database.Database.Delete(SuperUser{}, "ID = ?", superUser.ID).Error
	if err != nil {
		return err
	}
	return nil
}