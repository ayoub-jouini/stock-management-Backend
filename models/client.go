package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ID uint `gorm:"primary Key;autoIncrement" json:"id"`
	CompanyID uint `gorm:"primary Key"`
	Name string `gorm:"size:255;not null" json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	Avatar string `json:"avatar"`

	Compnay Company `gorm:"foreignKey:CompanyID;references:ID"`
}

type Clients []Client 

func (client Client) Save() (error) {
	err := database.Database.Create(&client).Error
	if err != nil {
		return err
	}
	return nil
}

func (client Client) FindById(id string) (error) {
	err := database.Database.Preload("Company").Where("ID=?", id).Find(&client).Error
	if err != nil {
		return err
	}
	return nil
}

func (clients Clients) FindAll(page string, limit string) (error) {
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	err := database.Database.Limit(intLimit).Offset(offset).Find(&clients).Error
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) Update() (error) {
	err := database.Database.Save(&client).Error
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) Delete() (error) {
	err := database.Database.Delete(Client{}, "ID = ?", client.ID).Error
	if err != nil {
		return err
	}
	return nil
}