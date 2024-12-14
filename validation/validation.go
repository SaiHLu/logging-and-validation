package validation

import (
	"context"
	"errors"
	"strings"

	"github.com/SaiHLu/logging-and-validation/util"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateParams(ctx context.Context, param interface{}) map[string]string {
	if err := validate.Struct(param); err != nil {
		return handleValidationErrors(err)
	}

	return nil
}

func ValidateQuery(ctx context.Context, query interface{}) map[string]string {
	if err := validate.Struct(query); err != nil {
		return handleValidationErrors(err)
	}

	return nil
}

func ValidateBody(ctx context.Context, body interface{}) map[string]string {
	if err := validate.Struct(body); err != nil {
		return handleValidationErrors(err)
	}

	return nil
}

func ValidateUpdateBody(ctx context.Context, body interface{}) map[string]string {
	validate.SetTagName("updatereq")

	if err := validate.Struct(body); err != nil {
		return handleValidationErrors(err)
	}

	return nil
}

func handleValidationErrors(err error) map[string]string {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		customErrorsFormat := make(map[string]string)
		for _, field := range validationErrors {
			customErrorsFormat[strings.ToLower(field.Field())] = util.FormatValidationMessage(field)
		}
		return customErrorsFormat
	}

	return nil
}
