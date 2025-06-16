package service

import (
	"crud/src/model"
	"crud/src/repository"
)

func CreateCompany(company models.Company) error {
	return repository.CreateCompany(&company)
}

func GetCompanies() ([]models.Company, error) {
	return repository.GetAllCompanies()
}

func GetCompany(id uint) (models.Company, error) {
	return repository.GetCompanyByID(id)
}

func UpdateCompany(company models.Company) error {
	return repository.UpdateCompany(&company)
}

func DeleteCompany(id uint) error {
	return repository.DeleteCompany(id)
}
