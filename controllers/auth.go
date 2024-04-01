package controllers

import (
	"diary_api/model"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Register(context *gin.Context) {
	var body model.User

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User {
		FirstName: body.FirstName,
		LastName: body.LastName,
		Email: body.Email,
		Phone: body.Phone,
		Password: body.Password,
		Avatar: body.Avatar,
		Role: body.Role
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}