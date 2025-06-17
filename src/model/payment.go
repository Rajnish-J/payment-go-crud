package models

import "time"

type Payment struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `json:"user_id"`
	ProductID    uint      `json:"product_id"`
	Amount       float64   `json:"amount"`
	Status       string    `json:"status"` // "Completed", "Failed"
	Mode         string    `json:"mode"`   // "Cash", "UPI", "Card"
	RewardTitle  string    `json:"reward_title"`
	RewardDesc   string    `json:"reward_desc"`
	PaymentTime  time.Time `json:"payment_time"`
	UseRewards   bool      `gorm:"-" json:"use_rewards"` // ignored in DB
}
