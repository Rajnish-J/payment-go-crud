package service

import (
	business "crud/src/business"
	"crud/src/config"
	models "crud/src/model"
)

func CreatePayment(req models.Payment) error {
	if err := business.ValidatePayment(req); err != nil {
		return err
	}
	payment := models.Payment{
		UserID:    req.UserID,
		Amount:    req.Amount,
		ProductID: req.ProductID,
	}
	result := config.DB.Create(&payment)
	return result.Error
}

func GetAllPayments() ([]models.Payment, error) {
	var payments []models.Payment
	result := config.DB.Find(&payments)
	return payments, result.Error
}

func GetPaymentByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	result := config.DB.First((&payment), id)
	if result.Error != nil {
		return &models.Payment{}, result.Error
	}
	return &payment, nil
}

func UpdatePayment(id uint, req models.Payment) error {
	if err := business.ValidatePayment(req); err != nil {
		return err
	}
	var payment models.Payment
	result := config.DB.First(&payment, id)
	if result.Error != nil {
		return result.Error
	}
	payment.ProductID = req.ProductID
	payment.Amount = req.Amount
	return config.DB.Save(&payment).Error
}

func DeletePayment(id uint) error {
	var payment models.Payment
	result := config.DB.First(&payment, id)
	if result.Error != nil {
		return result.Error
	}
	return config.DB.Delete(&payment).Error
}
