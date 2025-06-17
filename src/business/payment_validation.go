package business

import (
	models "crud/src/model"
	"errors"
)

func ValidatePayment(req models.Payment) error {
	if req.UserID == 0 {
		return errors.New("user_id is required")
	}
	if req.ProductID == 0 {
		return errors.New("product_id is required")
	}
	if req.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	return nil
}
