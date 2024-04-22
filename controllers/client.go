package controllers

import (
	"net/http"
	"stock_management/helper"
	"stock_management/models"

	"github.com/gin-gonic/gin"
)

func GetAllClients(contect *gin.Context) {
	var page string = context.DefaultQuery("page", "1")
	var limit string = context.DefaultQuery("limit", "10")

	var clients models.Clients
	if err := clients.FindAll(page, limit); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if the user is authorized to acces client info

	context.JSON(http.StatusOK, gin.H{"data": clients})
}

func GetClientById(context *gin.Context) {
	clientID := context.Param("id")

	var client models.Client
	if err := Client.FindById(clientID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if the user is authorized to acces client info

	context.JSON(http.StatusOK, gin.H{"data": client})
}

func AddClient(context *gin.Context) {
	var input models.Client
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error", err.Error()})
		return 
	}

	client := models.Client{
		CompanyID : user.CompanyID,
		Name : input.Name,
		Type : input.Type,
		Description : input.Description,
		Avatar : input.Avatar
	}

	if err = client.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": client})
}

func UpdateClient(context *gin.Context) {
	var input models.Client
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	clientID := context.Param("id")
	var client models.Client
	if err := client.FindById(clientID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	//check if user is authorized to update this client (role & company)
	
	client.Name = input.Name
	client.Type = input.Type
	client.Description = input.Description
	client.Avatar = input.Avatar

	if err := client.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": client})
}

func DeleteClient(context *gin.Context) {
	clientId := context.Param("id")

	var client models.Client
	if err := client.FindById(clientID); err := nil {
		context.JSON(http.StatusBadRequest, gin.H{"error", err.Error()})
		return 
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user is authorized to delete this client (role & company)

	if err := Client.Delete(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "deleted !"})
}

