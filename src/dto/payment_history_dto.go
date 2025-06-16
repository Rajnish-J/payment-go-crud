package dto

import "time"

type CreatePaymentHistoryRequest struct {
	UserID    uint      `json:"user_id" binding:"required"`
	ProductID uint      `json:"product_id" binding:"required"`
	Amount    float64   `json:"amount" binding:"required"`
	Timestamp time.Time `json:"timestamp"` // Optional: you can default this in service if needed
}
