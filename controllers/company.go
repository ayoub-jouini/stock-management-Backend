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

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var companies []models.Company

	err := Ctr.DB.Limit(intLimit).Offset(offset).Find(&companies).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": companies})
}

func (Ctr *CompanyControllers) GetCompanyByID(context *gin.Context) {
	companyID := context.Param("id")

	var company models.Company
	err := Ctr.DB.Preload("Employees").Preload("Admin").First(&company, "ID = ?", companyID).Error
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

	res := Ctr.DB.Create(&company)

	if res.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": company})
}