package dto

type CreateProductRequest struct{
	Name      string `json:"name" binding:"required`
	CompanyID uint `json:"company_id" binding:"required`
	Price     float64 `json:"price" binding:"required`
}