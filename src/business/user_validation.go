package business

import (
	"crud/src/model"
	"errors"
)

func ValidateUser(req models.User) error {
	if req.Name == "" || req.Email == "" || req.Password == "" || req.Age <= 0{
		return errors.New("all fields are required")
	}
	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
} 
