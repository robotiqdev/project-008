package numparse

import (
	"fmt"
	"strconv"

	calcerrors "github.com/robotiqdev/project-008/internal/errors"
)

// ParseNumber parses a string as a 64-bit floating point number.
// Returns (0, error) if the string cannot be parsed.
func ParseNumber(s string) (float64, error) {
	v, parseErr := strconv.ParseFloat(s, 64)
	if parseErr != nil {
		return 0, &calcerrors.CalcError{
			Code:    calcerrors.ErrCodeInvalidInput,
			Message: fmt.Sprintf("Error: invalid numeric input: %q", s),
			Wrapped: parseErr,
		}
	}
	return v, nil
}

// Validate checks whether s is a valid numeric string.
// Returns nil if valid, or an error if not.
func Validate(s string) error {
	_, err := ParseNumber(s)
	return err
}
