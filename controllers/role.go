package controllers

import (
	"net/http"
	"stock_management/helper"
	"stock_management/models"

	"github.com/gin-gonic/gin"
)

func GetAllRoles(context *gin.Context) {
	var page string = context.DefaultQuery("page", "1")
	var limit string = context.DefaultQuery("limit", "10")

	var roles models.Roles
	if err := roles.FindAll(page, limit); err != nil {
		context.JSON(http.StatusBadRequestn, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": roles})
}

func GetRoleById(context *gin.Context) {
	roleId := context.Param("id")

	var role models.Role
	if err := role.FindById(roleId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": role})
}

func AddRole(context *gin.Context) {
	var input models.Role
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//check if user is authorized

	role := models.Role{
		Name : input.Name
		Description input.Description
	}

	if err = role.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": role})
}

func UpdateRole(context *gin.context) {
	var input models.Role
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleId := context.Param("id")
	var role models.Role
	if err := role.FindById(roleId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user is authorized

	role.Name = input.Name
	role.Description = input.Description

	if err = role.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": role})
}

func DeleteRole(context *gin.Context) {
	roleId := context.Param("id")

	var role models.Role
	if err := role.FindById(roleId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user is authorized

	if err := role.Delete(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "deleted !"})
}