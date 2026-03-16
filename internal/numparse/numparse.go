package numparse

import (
	"fmt"
	"math"
	"strconv"

	calcerrors "github.com/robotiqdev/project-008/internal/errors"
)

// ParseNumber parses a string as a 64-bit floating point number.
// Returns (0, error) if the string cannot be parsed.
// The calculator only supports finite numbers: Inf and NaN are rejected even
// when strconv.ParseFloat would accept them (e.g. "Inf", "NaN", overflow values).
func ParseNumber(s string) (float64, error) {
	v, parseErr := strconv.ParseFloat(s, 64)
	if parseErr != nil {
		return 0, &calcerrors.CalcError{
			Code:    calcerrors.ErrCodeInvalidInput,
			Message: fmt.Sprintf("Error: invalid numeric input: %q", s),
			Wrapped: parseErr,
		}
	}
	// Reject non-finite results: Inf (from string literals or overflow) and NaN.
	// The calculator only operates on finite numbers; non-finite values represent
	// undefined states that must not propagate through the calculation engine.
	if math.IsInf(v, 0) || math.IsNaN(v) {
		return 0, &calcerrors.CalcError{
			Code:    calcerrors.ErrCodeInvalidInput,
			Message: fmt.Sprintf("Error: invalid numeric input: %q", s),
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
