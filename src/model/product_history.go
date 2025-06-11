package models

import "time"

type PaymentHistory struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
	Amount    float64
	Timestamp time.Time
}
