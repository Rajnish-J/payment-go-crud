package models

import "time"

type ProductHistory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	Action    string    `json:"action"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
	User      User      `gorm:"foreignKey:UserID"`
	Product   Product   `gorm:"foreignKey:ProductID"`
}