package dto

type CreateCompanyRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
