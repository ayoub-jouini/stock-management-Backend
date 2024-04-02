package controllers

import (
	"net/http"
	"stock_management/helper"
	"stock_management/models"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type AuthControllers struct {
	DB *gorm.DB
}

func AuthControllersInit(DB *gorm.DB) AuthControllers {
	return AuthControllers{DB}
}

func (ctr *AuthControllers) Register(context *gin.Context) {
	var body models.User

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User {
		FirstName: body.FirstName,
		LastName: body.LastName,
		Email: body.Email,
		Phone: body.Phone,
		Password: body.Password,
		Avatar: body.Avatar,
		// Role: body.Role
	}

	savedUser, err := ctr.DB.Create(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func (ctr *AuthControllers) Login(context *gin.Context) {
	var input models.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserByEmail(input.Email)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}