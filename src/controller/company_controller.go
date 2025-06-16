package controller

import (
	"crud/src/model"
	"crud/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateCompany(company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}
	c.JSON(http.StatusCreated, company)
}

func GetCompanies(c *gin.Context) {
	companies, err := service.GetCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch companies"})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func GetCompanyByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	company, err := service.GetCompany(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, company)
}

func UpdateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.UpdateCompany(company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
		return
	}
	c.JSON(http.StatusOK, company)
}

func DeleteCompany(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteCompany(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
