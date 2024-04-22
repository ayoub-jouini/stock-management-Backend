package controllers

import (
	"net/http"
	"stock_management/helper"
	"stock_management/models"

	"github.com/gin-gonic/gin"
)

func GetAllBills(context *gin.Context) {
	var page string = context.DefaultQuery("page", "1")
	var limit string = context.DefaultQuery("limit", "10")

	var bills models.Bills
	if err := bills.FindAll(page, limit); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": bills})
}

func GetBillByID(context *gin.Context) {
	billId := context.Param("id")

	var bill models.Bill
	if err := bill.FindById(billId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": bill})
}

func AddBill(context *gin.Context) {
	var input models.Bill
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//check if user is authorized

	bill := models.Bill{
		BillNum : input.DocNum
		CompanyID : user.CompanyID
		ClientID : user.ClientID
		DueDate : input.DueDate
		Product : input.Product
		Service : input.Service
	}

	if err = bill.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": bill})
}

func UpdateBill(context *gin.Context) {
	var input models.Bill
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	billId := context.Param("id")
	var bill models.Bill
	if err := bill.FindById(billId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user is authorized

	bill.BillNum = input.DocNum
	bill.ClientID = input.ClientID
	bill.DueDate = input.DueDate
	bill.Product = input.Product
	bill.Service = input.Service

	if err = bill.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": bill})
}

func DeleteBill(context *gin.Context) {
	billId := context.Param("id")

	var bill models.Bill
	if err := bill.FindById(billId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user is authorized

	if err := bill.Delete(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "deleted !"})
}