package models

import "time"

type PaymentHistory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Amount    float64   `json:"amount"`
	ProductID uint      `json:"product_id"`
	TimeStamp time.Time `json:"TimeStamp"`
}
