package errors

import "fmt"

// NullArgumentError represents a error that is used to signal that a provided argument is null
type NullArgumentError struct {
	ArgumentName string
}

// Error returns the string representation of the error
func (e *NullArgumentError) Error() string {
	return fmt.Sprintf("Parameter %s can't be null", e.ArgumentName)
}

// NewNullArgumentError returns a preformatted error for null arguments
func NewNullArgumentError(argumentName string) *NullArgumentError {
	return &NullArgumentError{argumentName}
}
