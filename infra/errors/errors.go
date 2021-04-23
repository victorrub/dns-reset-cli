package errors

import (
	"runtime"
	"strings"
)

// errorWrapper defines the interface for an error wrapper that extends an error with additional information
type errorWrapper interface {
	Error() string
	GetOriginalError() error
}

// Wrap wraps an error with a context message and adds execution path
func Wrap(err error, messages ...string) error {
	if err != nil {
		// get caller function path
		pc := make([]uintptr, 10)
		runtime.Callers(2, pc)
		funcRef := runtime.FuncForPC(pc[0])

		pathArr := strings.Split(funcRef.Name(), "/")

		path := pathArr[len(pathArr)-1]

		return &WrappedError{
			originalError: err,
			path:          path,
			messages:      messages,
		}
	}

	return nil
}

// GetOriginalError returns the original error if the provided error is a WrappedError.
// Returns the provided error otherwise
func GetOriginalError(err error) error {
	wrappedErr, ok := err.(errorWrapper)
	if ok {
		return wrappedErr.GetOriginalError()
	}

	return err
}
