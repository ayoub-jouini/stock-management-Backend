package controllers

import (
	"net/http"
	"stock_management/helper"
	"stock_management/models"

	"github.com/gin-gonic/gin"
)

func GetAllTaxes(context *gin.Context) {
	var page string = context.DefaultQuery("page", "1")
	var limit string = context.DefaultQuery("limit", "10")

	var taxes models.Taxes
	if err := taxes.FindAll(page, limit); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": taxes})
}

func GetTaxByID(context *gin.Context) {
	taxId := context.Param("id")

	var tax models.Tax
	if err := tax.FindById(taxId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": tax})
}

func AddTax(context *gin.Context) {
	var input models.Tax
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//check if user is authorized

	tax := models.Tax{
		Name : input.Name
		Type : input.Type
		Value : input.Value
		CompanyID : user.CompanyID
	}

	if err = tax.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": tax})
}

func UpdateTax(context *gin.Context) {
	var input models.Tax
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taxId := context.Param("id")
	var tax models.Tax
	if err := tax.FindById(taxId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user is authorized

	tax.Name = input.Name
	tax.Type = input.Type
	tax.Value = input.Value

	if err = bill.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": bill})
}

func DeleteTax(context *gin.Context) {
	taxId := context.Param("id")

	var tax models.Bill
	if err := tax.FindById(taxId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user is authorized

	if err := tax.Delete(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "deleted !"})
}