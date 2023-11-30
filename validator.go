package rvalidator

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// getErrorMessage retrieves a custom error message for a specific validation error.
// Parameters:
//   - fe validator.FieldError: The validation error object from the validator package.
//   - s interface{}: The struct instance being validated. It's of type interface{}
//     to allow for any struct type.
//
// Returns:
// - A string representing the custom error message.
// Description:
// - Uses reflection to inspect the struct and find the field that caused the validation error.
// - Extracts the custom error message for the field from its 'errormessage' tag.
// - If no custom message is defined, it defaults to the standard error message from the validator.
func getErrorMessage(fe validator.FieldError, s interface{}) string {
	val := reflect.ValueOf(s)
	field, ok := val.Type().FieldByName(fe.Field())

	if !ok {
		return "Unknown field"
	}

	errorMessage := field.Tag.Get("errormessage")
	if errorMessage == "" {
		return fe.Error()
	}

	return errorMessage
}

// Validate performs validation on a struct and returns a slice of custom error messages.
// Parameters:
// - s interface{}: The struct instance to validate. Accepts any type due to the interface{} signature.
// Returns:
// - A slice of strings, each representing a custom error message for a validation failure.
// - An error object, which is non-nil if there are non-validation errors (like an invalid input type).
// Description:
// - First checks if the input is a struct. Returns an error if not.
// - Creates a new validator instance and performs validation on the provided struct.
// - If an InvalidValidationError is encountered, it is returned immediately.
// - For validation errors, iterates over them, using getErrorMessage to gather custom error messages.
// - Returns a slice of custom error messages and nil if there are only validation errors.
// Usage:
// - Users should call this function to validate their structs. It simplifies handling validation errors and presents them clearly.
func Validate(s interface{}) ([]string, error) {
	// Ensure the input is a struct
	if reflect.ValueOf(s).Kind() != reflect.Struct {
		return nil, errors.New("input must be a struct")
	}

	var errorsStrings []string
	valid := validator.New()
	err := valid.Struct(s)
	if err != nil {
		// Handle non-validation errors
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err
		}

		// Handle validation errors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range validationErrors {
				errorsStrings = append(errorsStrings, getErrorMessage(err, s))
			}
		}
		return errorsStrings, nil
	}

	return nil, nil
}
