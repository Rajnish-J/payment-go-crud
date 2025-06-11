package models

type Product struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CompanyID uint
	Price     float64
}
