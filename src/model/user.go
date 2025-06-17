package models

type User struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password,omitempty"`
	Age      int     `json:"age"`
	Address  string  `json:"address"`
	Phone    string  `json:"phone"`
	Payments []Payment `gorm:"foreignKey:UserID" json:"payments,omitempty"`
}
