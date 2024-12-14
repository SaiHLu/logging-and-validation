package util

import "github.com/go-playground/validator/v10"

func FormatValidationMessage(err validator.FieldError) string {
	field, tag, param := err.Field(), err.Tag(), err.Param()
	switch tag {
	case "required":
		return field + " field is required"
	case "email":
		return "Invalid email"
	case "min":
		return field + " should have minimum length of " + param
	case "max":
		return field + " should have maximum length of " + param
	default:
		return field + " is invalid"
	}
}
