package models

type Company struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Address string
}
