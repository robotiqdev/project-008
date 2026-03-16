package errors

import "errors"

// ErrInvalidInput is the sentinel error for invalid numeric input.
var ErrInvalidInput = errors.New("invalid input")

// ErrCodeInvalidInput is the error code for invalid input errors.
const ErrCodeInvalidInput = "INVALID_INPUT"

// CalcError is a structured error type for calculator errors.
type CalcError struct {
	Code    string
	Message string
	Wrapped error
}

// Error implements the error interface.
func (e *CalcError) Error() string {
	return e.Message
}

// Unwrap returns the wrapped error for errors.Is/As traversal.
func (e *CalcError) Unwrap() error {
	return e.Wrapped
}

// Is allows errors.Is(err, ErrInvalidInput) to match based on the Code field.
func (e *CalcError) Is(target error) bool {
	if target == ErrInvalidInput {
		return e.Code == ErrCodeInvalidInput
	}
	return false
}
