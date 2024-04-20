package controllers

import (
	"fmt"
	"net/http"
	"stock_management/helper"
	"stock_management/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) {
	var page string = context.DefaultQuery("page", "1")
	var limit string = context.DefaultQuery("limit", "10")

	users, err := models.FindAllUsers(&page, &limit)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserByID(context *gin.Context) {
	userID := context.Param("id")

	user, err := models.FindUserById(userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUserById(context *gin.Context) {
	userID := context.Param("id")
	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if fmt.Sprint(user.ID) != userID {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Not current user"})
		return
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Phone = input.Phone
	user.Password = input.Password
	user.Avatar = input.Avatar
	user.Role = input.Role

	if err = user.UpdateUser(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": user})
}

// func DeleteUserById(context *gin.Context) {
// 	userID := gin.Param("id")

// 	user, err := helper.CurrentUser(context)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if user.ID != company.Admin {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "Not Admin"})
// 		return
// 	}
	
// 	if err := Model.DeleteCompany(compnayID); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	context.JSON(http.StatusCreated, gin.H{"message": "deleted !"})
// }