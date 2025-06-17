package repository

import (
	"crud/src/model"
	"crud/src/config"
)

func CreatePayment(p *models.Payment) error {
	return config.DB.Create(p).Error
}

func GetAllPayments() ([]models.Payment, error) {
	var payments []models.Payment
	err := config.DB.Find(&payments).Error
	return payments, err
}

func GetPaymentByID(id uint) (*models.Payment, error) {
	var p models.Payment
	err := config.DB.First(&p, id).Error
	return &p, err
}

func UpdatePayment(p *models.Payment) error {
	return config.DB.Save(p).Error
}

func DeletePayment(id uint) error {
	return config.DB.Delete(&models.Payment{}, id).Error
}
