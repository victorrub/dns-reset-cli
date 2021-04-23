package errors

import (
	"fmt"
	"strings"
)

// ValidationError represents an input validation error
type ValidationError struct {
	Field   string            `json:"field_name,omitempty"`
	Message string            `json:"message,omitempty"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

func (e *ValidationError) Error() string {
	var errorList []string

	for _, err := range e.Errors {
		errorList = append(errorList, err.Error())
	}

	output := e.Message

	if len(errorList) > 0 {
		output += fmt.Sprintf("\n - %v", strings.Join(errorList, ";\n - "))
	}

	return fmt.Sprintf("%v: %v", e.Field, output)
}

// AddError adds a new validation error to the chain
func (e *ValidationError) AddError(field string, message string) {
	e.Errors = append(e.Errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

// NewValidationError returns a ValidationError instance with the provided parameters
func NewValidationError(field string, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}
