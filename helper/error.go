package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ErrorPanic(err error) error {
	if err != nil {
		// panic(fmt.Sprintf("Panic Error: %d", err))
		return err
	}
	return nil
}

func FormatValidationErrors(errs validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, fieldErr := range errs {
		fieldName := fieldErr.Field()
		errors[fieldName] = fmt.Sprintf(
			"Validation failed for '%s': %s (expected %s)",
			fieldName,
			fieldErr.Value(),
			fieldErr.Tag(),
		)
	}
	return errors
}
