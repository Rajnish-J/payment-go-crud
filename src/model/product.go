package models

type Product struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CompanyID uint
	Company   Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price     float64
}
