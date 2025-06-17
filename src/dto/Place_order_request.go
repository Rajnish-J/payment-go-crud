// dto/place_order.go
package dto

type PlaceOrderRequest struct {
	UserID       uint    `json:"user_id"`
	ProductID    uint    `json:"product_id"`
	PaymentMode  string  `json:"payment_mode"` // e.g. "UPI", "Card"
	UseRewards   bool    `json:"use_rewards"`
}
