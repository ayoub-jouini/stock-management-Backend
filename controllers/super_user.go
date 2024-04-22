package controllers

import (
	"net/http"
	"stock_management/helper"
	"stock_management/models"

	"github.com/gin-gonic/gin"
)

func GetAllSuperUser(context *gin.Context) {
	var page string = context.DefaultQuery("page", "1")
	var limit string = context.DefaultQuery("limit", "10")

	var superUsers models.SuperUsers 
	if err := superUsers.FindAll(page, limit); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": superUsers})
}

func GetBillByID(context *gin.Context) {
	superUserId := context.Param("id")

	var superUser models.SuperUser
	if err := superUser.FindById(superUserId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": superUser})
}

func AddSuperUser(context *gin.Context) {
	var input models.SuperUser
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//check if user is authorized

	superUser := models.SuperUser{
		UserID : input.UserID
	}

	//check if user exist

	if err = superUser.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": superUser})
}

func DeleteSuperUser(context *gin.Context) {
	superUserId := context.Param("id")

	var superUser models.SuperUser
	if err := superUser.FindById(superUserId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user is authorized

	if err := superUser.Delete(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "deleted !"})
}