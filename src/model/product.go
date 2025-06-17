package models

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	CompanyID   uint      `json:"company_id"`
	Company     Company   `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"company"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Payments    []Payment `gorm:"foreignKey:ProductID" json:"payments,omitempty"`
}
