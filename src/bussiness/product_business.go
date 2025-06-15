package business

import (
	"crud/src/dto"
	models "crud/src/model"
	"crud/src/repository"
	"errors"
)
var repo repository.ProductRepo

func GetAllProduts() ([]models.Product,error){
	return repo.GetAllProducts()
}

func GetProductByID(id uint) (models.Product, error) {
	return repo.GetProductByID(id)
}

func CreateProduct(req dto.CreateProductRequest) error {
	if req.Name == "" || req.CompanyID <=0 || req.Price <=0 {
		return errors.New("all fields are required")
	}
	product:= models.Product{
		Name: req.Name,
		CompanyID: req.CompanyID,
		Price: req.Price,
	}
	repo.CreateProduct(product)
	return nil
}


func UpdateProduct(id uint, req dto.CreateProductRequest) error {
	if req.Name == "" || req.CompanyID <= 0 || req.Price <= 0 {
		return errors.New("all fields are required")
	}
	product := models.Product{
		ID:        id,
		Name:      req.Name,
		CompanyID: req.CompanyID,
		Price:     req.Price,
	}
	return repo.UpdateProduct(id, product)
}

func DeleteProduct(id uint) error {
	return repo.DeleteProduct(id)
}