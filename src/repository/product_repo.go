package repository

import (
	"crud/src/config"
	models "crud/src/model"
)

type ProductRepo struct {
}

func (repo *ProductRepo) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := config.DB.Find(&products)
	return products, result.Error
}

func (repo *ProductRepo) GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

func (repo *ProductRepo) CreateProduct(product models.Product) error {
	return config.DB.Create(&product).Error
}


func (repo *ProductRepo) UpdateProduct(id uint, product models.Product) error {
	var existProduct models.Product
	result := config.DB.First(&existProduct, id)
	if result.Error != nil {
		return result.Error
	}

	// Update fields
	existProduct.Name = product.Name
	existProduct.CompanyID = product.CompanyID
	existProduct.Price = product.Price

	return config.DB.Save(&existProduct).Error
}


func (repo *ProductRepo) DeleteProduct(id uint) error {
	var existProduct models.Product
	result := config.DB.First(&existProduct, id)
	if result.Error != nil {
		return result.Error
	}
	return config.DB.Delete(&existProduct).Error
}
