package repository

import models "crud/src/model"

type Irepo interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (models.Product, error)
	CreateProduct(product models.Product) error
	UpdateProduct(id uint, product models.Product) error
	DeleteProduct(id uint) error
}