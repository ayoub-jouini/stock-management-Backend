package controllers

import (
    "diary_api/helper"
    "diary_api/model"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GellAllCompanies(context *gin.Context) {}

func AddCompany(context *gin.Context) {
	var company model.Company
	if err := context.ShouldBindJSON(&company); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company := model.Company {
		RegNum: body.RegNum,
		Name: body.Name,
		Description: body.Description,
		Email: body.Email,
		Address: body.Address,
		City: body.City,
		Country: body.Country
		Phone: body.Phone
		Logo: body.Logo
	}

	savedCompany, err := company.save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedCompany})
}