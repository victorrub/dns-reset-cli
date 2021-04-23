package domain

type errorString string

// Error returns the error message
func (err errorString) Error() string {
	return string(err)
}

// ErrNotFound represents a not found error
const ErrNotFound = errorString("Not found")

// ErrCurrentLocation .
const ErrCurrentLocation = errorString("It was not possible to set the requested location. The current network location is %s. Try again.")
