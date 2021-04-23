package errors

import "fmt"

// WrappedError holds an error wrapped with a context message
type WrappedError struct {
	originalError error
	path          string
	messages      []string
}

// Error returns the string representation of the error
func (err WrappedError) Error() string {
	if len(err.messages) > 0 {
		retVal := fmt.Sprintf("%s: ", err.path)

		for _, message := range err.messages {
			retVal += message + "; "
		}

		return fmt.Sprintf("%s => %v", retVal, err.originalError)
	}

	return fmt.Sprintf("%s => %v", err.path, err.originalError)
}

// GetOriginalError returns the original error
func (err WrappedError) GetOriginalError() error {
	if err.originalError != nil {
		originalError, ok := (err.originalError).(errorWrapper)
		if ok {
			return originalError.GetOriginalError()
		}
	}

	return err.originalError
}
