package models

type Company struct {
	ID      uint      `gorm:"primaryKey" json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Products []Product `gorm:"foreignKey:CompanyID" json:"products,omitempty"`
}
