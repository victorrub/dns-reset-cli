package errors

// ApplicationError represents a common application error structure
type ApplicationError struct {
	Message string
	Path    string
}

// Error returns the string representation of the error
func (e *ApplicationError) Error() string {
	return e.Message
}

// NewApplicationError returns a ApplicationError instance
func NewApplicationError(path string, message string) *ApplicationError {
	return &ApplicationError{
		Message: message,
		Path:    path,
	}
}
