package business

import (
	"crud/src/model"
	"errors"
	"strings"
	"unicode"
)

// ValidateCompany checks if the company data is valid
func ValidateCompany(req models.Company) error {
	// Trim spaces
	name := strings.TrimSpace(req.Name)
	address := strings.TrimSpace(req.Address)

	if name == "" || address == "" {
		return errors.New("name and address fields are required")
	}

	if len(name) < 3 {
		return errors.New("company name must be at least 3 characters")
	}

	if len(name) > 100 {
		return errors.New("company name must not exceed 100 characters")
	}

	if len(address) < 5 {
		return errors.New("address must be at least 5 characters")
	}

	if hasSpecialChars(name) {
		return errors.New("company name must not contain special characters")
	}

	return nil
}

// hasSpecialChars checks for special characters in the name
func hasSpecialChars(s string) bool {
	for _, char := range s {
		if !(unicode.IsLetter(char) || unicode.IsSpace(char)) {
			return true
		}
	}
	return false
}
