package repository

import (
	"crud/src/config"
	"crud/src/model"
)

func CreateCompany(company *models.Company) error {
	return config.DB.Create(company).Error
}

func GetAllCompanies() ([]models.Company, error) {
	var companies []models.Company
	err := config.DB.Find(&companies).Error
	return companies, err
}

func GetCompanyByID(id uint) (models.Company, error) {
	var company models.Company
	err := config.DB.First(&company, id).Error
	return company, err
}

func UpdateCompany(company *models.Company) error {
	return config.DB.Save(company).Error
}

func DeleteCompany(id uint) error {
	return config.DB.Delete(&models.Company{}, id).Error
}
