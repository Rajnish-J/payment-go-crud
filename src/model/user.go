package models

type User struct {
	// gorm.Model     // This will add fields ID, CreatedAt, UpdatedAt, DeletedAt
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Age      int    `json:"age"`
}
