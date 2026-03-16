package numparse

import "strconv"

// ParseNumber parses a string as a 64-bit floating point number.
// Returns (0, error) if the string cannot be parsed.
func ParseNumber(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

// Validate checks whether s is a valid numeric string.
// Returns nil if valid, or an error if not.
func Validate(s string) error {
	_, err := ParseNumber(s)
	return err
}
