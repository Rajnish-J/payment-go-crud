package service

import (
	"crud/src/business"
	"crud/src/config"
	models "crud/src/model"
	"crud/src/utils"
	"errors"
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

func PlaceOrderService(req models.Payment) error {
	tx := config.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := business.PlaceOrder(tx, req); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func LoginUser(req models.LoginRequest) (string, error) {
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return "", errors.New("invalid email or user not found")
	}

	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	// 	return "", errors.New("invalid password")
	// }

	if user.Password != req.Password {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(user.ID, "user") // Replace "user" with user.Role if you use roles
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
