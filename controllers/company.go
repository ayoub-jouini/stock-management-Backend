package controllers

import (
	"net/http"
	"stock_management/models"

	"github.com/gin-gonic/gin"
)

func GellAllCompanies(context *gin.Context) {}

func AddCompany(context *gin.Context) {
	var body models.Company
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// user, err := helper.CurrentUser(context)

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	body = models.Company {
		RegNum: body.RegNum,
		Name: body.Name,
		Description: body.Description,
		Email: body.Email,
		Address: body.Address,
		City: body.City,
		Country: body.Country,
		Phone: body.Phone,
		Logo: body.Logo,
	}

	savedCompany, err := body.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedCompany})
}