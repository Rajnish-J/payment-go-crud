package repository

import (
	"crud/src/model"
	"crud/src/config"
)

func CreatePayment(p *models.PaymentHistory) error {
	return config.DB.Create(p).Error
}

func GetAllPayments() ([]models.PaymentHistory, error) {
	var payments []models.PaymentHistory
	err := config.DB.Find(&payments).Error
	return payments, err
}

func GetPaymentByID(id uint) (*models.PaymentHistory, error) {
	var p models.PaymentHistory
	err := config.DB.First(&p, id).Error
	return &p, err
}

func UpdatePayment(p *models.PaymentHistory) error {
	return config.DB.Save(p).Error
}

func DeletePayment(id uint) error {
	return config.DB.Delete(&models.PaymentHistory{}, id).Error
}
