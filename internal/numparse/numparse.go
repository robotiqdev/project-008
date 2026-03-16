package numparse

import (
	"strconv"

	calcerrors "github.com/repo/calculator/internal/errors"
)

// ParseNumber parses s as a float64. Returns ErrInvalidInput on failure.
func ParseNumber(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, calcerrors.ErrInvalidInput
	}
	return v, nil
}
