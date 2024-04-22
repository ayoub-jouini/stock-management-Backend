package controllers

import (
	"net/http"
	"stock_management/models"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(context *gin.Context) {
	var page string = context.DefaultQuery("page", "1")
	var limit string = context.DefaultQuery("limit", "10")

	var categories models.Categories
	if err := categories.FindAll(page, limit); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": categories})
}

func GetCategoryById(context *gin.Context) {
	catgoryID := context.Param("id")
	
	var category models.Category
	if err := category.FindById(catgoryID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": category})
}

func AddCategory(context *gin.Context) {
	var input models.Category
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	//get Company

	category := models.Category{
		Name: input.Name,
		Description: input.Description,
		CompanyID: input.CompanyID,
	}

	if err := category.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data" : category})
}

func UpdateCategory(context *gin.Context) {
	var input models.Category
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryID := context.Param("id")
	var category models.Category
	if err := category.FindById(categoryID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// get company && check user

	category.Name = input.Name
	category.Description = input.Description

}

func DeleteCategory(context *gin.Context) {
	var categoryId = context.Param("id")

	var category models.Category
	if err := category.FindById(categoryId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	//check user company and category company

	if err := category.Delete(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted!"})
}