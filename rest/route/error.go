package route

import (
	"fmt"

	"github.com/pkg/errors"
)

// BadJSONError is an error that occurs when the JSON is invalid.
type BadJSONError struct {
	Message string
}

// NewBadJSONError creates a new BadJSONError.
func NewBadJSONError(format string, args ...any) BadJSONError {
	return BadJSONError{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error implements the error interface.
func (e BadJSONError) Error() string {
	return e.Message
}

// IsBadJSONError checks if the error is a BadJSONError.
func IsBadJSONError(err error) bool {
	var target BadJSONError

	return errors.As(err, &target)
}
