package business

import (
	"crud/src/dto"
	"errors"
)

func ValidateUser(req dto.CreateUserRequest) error {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return errors.New("all fields are required")
	}
	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}
