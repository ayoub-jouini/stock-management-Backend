package controllers

import (
	"net/http"
	"stock_management/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompanyControllers struct {
	DB *gorm.DB
}

func CompanyControllersInit(DB *gorm.DB) CompanyControllers {
	return CompanyControllers{DB}
}

func (Ctr *CompanyControllers) GetAllCompanies(context *gin.Context) {
	var page = context.DefaultQuery("page", "1")
	var limit = context.DefaultQuery("limit", "10")

	companies, err := models.FindAllCompanies(page, limit)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": companies})
}

func (Ctr *CompanyControllers) GetCompanyByID(context *gin.Context) {
	companyID := context.Param("id")

	company, err := models.FindCompanyByID(uint(companyID))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": company})
}

func (Ctr *CompanyControllers) AddCompany(context *gin.Context) {
	var input models.Company
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
	}

	res, err := company.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": company})
}