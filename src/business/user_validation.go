package business

import (
	models "crud/src/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

func ValidateUser(req models.User) error {
	if req.Name == "" || req.Email == "" || req.Password == "" || req.Age <= 0 {
		return errors.New("all fields are required")
	}
	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}

func PlaceOrder(tx *gorm.DB, req models.Payment) error {
	// Check if user exists
	var user models.User
	if err := tx.First(&user, req.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	// Check if product exists
	var product models.Product
	if err := tx.First(&product, req.ProductID).Error; err != nil {
		return errors.New("product not found")
	}

	finalAmount := product.Price
	rewardTitle := ""
	rewardDesc := ""

	// Apply dummy reward logic
	if req.UseRewards {
		finalAmount -= 10 // flat reward
		rewardTitle = "₹10 OFF"
		rewardDesc = "Used available rewards"
	}

	// Save payment
	payment := models.Payment{
		UserID:      req.UserID,
		ProductID:   req.ProductID,
		Amount:      finalAmount,
		Status:      "Completed",
		Mode:        req.Mode,
		RewardTitle: rewardTitle,
		RewardDesc:  rewardDesc,
		PaymentTime: time.Now(),
	}

	if err := tx.Create(&payment).Error; err != nil {
		return err
	}

	// Add to product history
	history := models.ProductHistory{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Timestamp: time.Now(),
	}
	if err := tx.Create(&history).Error; err != nil {
		return err
	}

	return nil
}
