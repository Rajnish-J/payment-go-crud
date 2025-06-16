package service

import (
	business "crud/src/business"
	"crud/src/dto"
	models "crud/src/model"
)

func GetAllProduts() ([]models.Product, error) {
	result, error := business.GetAllProduts()
	return result, error
}

func GetProductByID(id uint) (models.Product, error) {
	product, err := business.GetProductByID(id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func CreateProduct(product dto.CreateProductRequest) error {
	err := business.CreateProduct(product)

	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(id uint,product dto.CreateProductRequest) error {
	err := business.UpdateProduct(id,product)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id uint) error {
	err := business.DeleteProduct(id)

	if err != nil {
		return err
	}
	return nil
}
