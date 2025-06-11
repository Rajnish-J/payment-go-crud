package repository

import (
	"crud/src/config"
	"crud/src/model"
)

func CreateUser(user models.User) error {
	return config.DB.Create(&user).Error
}
