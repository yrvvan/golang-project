package validators

import (
	"strings"

	"golang-project/models"
	"golang-project/database"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			errors[strings.ToLower(fe.Field())] = validationErrorMessage(fe)
		}
	}

	return errors
}

func validationErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters"
	case "max":
		return fe.Field() + " must be no more than " + fe.Param() + " characters"
	default:
		return fe.Field() + " is invalid"
	}
}

func EmailNotExist(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err == nil {
		// User exists, return false (email already taken)
		return false
	}

	// Email does not exist, return true
	return true
}
