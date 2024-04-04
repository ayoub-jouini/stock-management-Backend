package controllers

import (
	"net/http"
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