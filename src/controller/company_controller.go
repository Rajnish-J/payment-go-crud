package controller

import (
	models "crud/src/model"
	"crud/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCompany godoc
// @Summary      Create a new company
// @Description  Adds a new company to the system
// @Tags         company
// @Accept       json
// @Produce      json
// @Param        company  body      models.Company  true  "Company to create"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /api/company [post]
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

// GetCompanyByID godoc
// @Summary      Get company by ID
// @Description  Retrieves a company by their ID
// @Tags         company
// @Produce      json
// @Param        id   path      int  true  "Company ID"
// @Success      200  {object}  models.Company
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/company/{id} [get]
func GetCompanies(c *gin.Context) {
	companies, err := service.GetCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch companies"})
		return
	}
	c.JSON(http.StatusOK, companies)
}

// GetAllCompanies godoc
// @Summary      Get all companies
// @Description  Retrieves all companys from the database
// @Tags         company
// @Produce      json
// @Success      200  {array}  models.Company
// @Router       /api/company [get]
func GetCompanyByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	company, err := service.GetCompany(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, company)
}

// UpdateCompany godoc
// @Summary      Update an existing company
// @Description  Updates company data by ID
// @Tags         company
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "Company ID"
// @Param        company  body      models.Company  true  "Updated company"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /api/company/{id} [put]
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

// DeleteCompany godoc
// @Summary      Delete a company
// @Description  Removes a company from the database by ID
// @Tags         company
// @Produce      json
// @Param        id   path      int  true  "Company ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/company/{id} [delete]
func DeleteCompany(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteCompany(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
