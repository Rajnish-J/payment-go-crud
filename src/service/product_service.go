package service

import (
	"crud/src/bussiness"
	"crud/src/dto"
) 

func GetAllProduts() ([]dto.CreateProductRequest, error) {
	result, error := business.GetAllProduts()
	return result, error
}

func GetProductByID(id uint) (dto.CreateProductRequest, error) {
	product, err := business.GetProductByID(id)
	if err != nil{
		return product, err
	}
	return product,nil
}

func CreateProduct(product dto.CreateProductRequest) error {
	err := business.CreateProduct(product)

	if err != nil{
		return err
	}
	return nil
}

func UpdateProduct(product dto.CreateProductRequest) error {
	err := business.UpdateProduct(product)

	if err != nil{
		return err
	}
	return nil
}

func DeleteProduct(id uint) error {
	err := business.DeleteProduct(id)

	 if err != nil{
		return err
	 }
	 return nil
}