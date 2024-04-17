package controllers

import (
	"net/http"
	"stock_management/helper"
	"stock_management/models"

	"github.com/gin-gonic/gin"
)

func GetAllCompanies(context *gin.Context) {
	var page string = context.DefaultQuery("page", "1")
	var limit string = context.DefaultQuery("limit", "10")

	companies, err := models.FindAllCompanies(&page, &limit)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": companies})
}

func GetCompanyByID(context *gin.Context) {
	companyID := context.Param("id")

	company, err := models.FindCompanyByID(companyID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": company})
}

func AddCompany(context *gin.Context) {
	var input models.Company
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	company := models.Company {
		RegNum: input.RegNum,
		Name: input.Name,
		Description: input.Description,
		Email: input.Email,
		Address: input.Address,
		City: input.City,
		Country: input.Country,
		Phone: input.Phone,
		Logo: input.Logo,
		Admin: user.ID,
	}

	_, err = company.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": company})
}

func UpdateCompanyById(context *gin.Context) {
	var input models.Company
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyID := gin.Param("id")
	company, err := models.FindCompanyByID(companyID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.ID != company.Admin {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Not Admin"})
		return
	}

	company.RegNum = input.RegNum,
	company.Name = input.Name,
	company.Description = input.Description,
	company.Email = input.Email,
	company.Address = input.Address,
	company.City = input.City,
	company.Country = input.Country,
	company.Phone = input.Phone,
	company.Logo = input.Logo,

	if err = company.UpdateCompany(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": company})
}

func DeleteCompanyById(context *gin.Context) {
	companyID := gin.Param("id")

	company, err := models.FindCompanyByID(companyID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.ID != company.Admin {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Not Admin"})
		return
	}
	
	if err := Model.DeleteCompany(compnayID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "deleted !"})
}