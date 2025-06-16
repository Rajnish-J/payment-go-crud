package service

import (
	"crud/src/model"
	"crud/src/repository"
)

func CreatePayment(p models.PaymentHistory) error {
	return repository.CreatePayment(&p)
}

func GetAllPayments() ([]models.PaymentHistory, error) {
	return repository.GetAllPayments()
}

func GetPaymentByID(id uint) (*models.PaymentHistory, error) {
	return repository.GetPaymentByID(id)
}

func UpdatePayment(p models.PaymentHistory) error {
	return repository.UpdatePayment(&p)
}

func DeletePayment(id uint) error {
	return repository.DeletePayment(id)
}
