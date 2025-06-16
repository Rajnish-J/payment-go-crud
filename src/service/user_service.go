package service

import (
	"crud/src/bussiness"
	"crud/src/config"
	"crud/src/model"
) 

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func CreateUser(req models.User) error {
	if err := business.ValidateUser(req); err != nil {
		return err
	}
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	result := config.DB.Create(&user)
	return result.Error
}

func UpdateUser(id uint, req models.User) error {
	if err := business.ValidateUser(req); err != nil {
		return err
	}
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.Name = req.Name
	user.Email = req.Email
	user.Password = req.Password
	return config.DB.Save(&user).Error
}

func DeleteUser(id uint) error {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return config.DB.Delete(&user).Error
}