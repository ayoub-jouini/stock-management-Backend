package controllers

import (
	"net/http"
	"stock_management/helper"
	"stock_management/models"

	"github.com/gin-gonic/gin"
)

func GetAllEstimationDocs(context *gin.Context) {
	var page string = context.DefaultQuery("page", "1")
	var limit string = context.DefaultQuery("limit", "10")

	var estimationDocs models.EstimationDocs
	if err := estimationDocs.FindAll(page, limit); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": estimationDocs})
}

func GetEstimationDocByID(context *gin.Context) {
	estimationDocId := context.Param("id")

	var estimationDoc models.EstimationDoc
	if err := estimationDoc.FindById(estimationDocId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": estimationDoc})
}

func AddEstimationDoc(context *gin.Context) {
	var input models.EstimationDoc
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//check if user is authorized

	estimationDoc := models.EstimationDoc{
		DocNum : input.DocNum
		CompanyID : user.CompanyID
		ClientID : user.ClientID
		DueDate : input.DueDate
		Product : input.Product
		Service : input.Service
	}

	if err = estimationDoc.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": estimationDoc})
}

func UpdateEstimationDoc(context *gin.Context) {
	var input models.EstimationDoc
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	estimationDocId := context.Param("id")
	var estimationDoc models.EstimationDoc
	if err := estimationDoc.FindById(estimationDocId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user is authorized
	
	estimationDoc.DocNum = input.DocNum
	estimationDoc.ClientID = input.ClientID
	estimationDoc.DueDate = input.DueDate
	estimationDoc.Product = input.Product
	estimationDoc.Service = input.Service

	if err = estimationDoc.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": estimationDoc})
}

func DeleteEstimationDoc(context *gin.Context) {
	estimationDocId := context.Param("id")

	var estimationDoc models.EstimationDoc
	if err := estimationDoc.FindById(estimationDocId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user is authorized

	if err := estimationDoc.Delete(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "deleted !"})
}