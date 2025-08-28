package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return fmt.Sprintf("must be at least %s characters", fe.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters", fe.Param())
	default:
		return fmt.Sprintf("failed on '%s' validation", fe.Tag())
	}
}

var Validate = validator.New()

func ValidateStruct(s interface{}) []string {
	err := Validate.Struct(s)

	if err == nil {
		return nil
	}

	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		// contoh: "Password: must be at least 6 characters"
		errors = append(errors, fmt.Sprintf("%s: %s", e.Field(), msgForTag(e)))
	}

	return errors
}
