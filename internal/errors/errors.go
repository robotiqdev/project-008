package errors

import (
	"fmt"
	"io"
)

// ErrCode identifies the category of a calculator error.
type ErrCode int

const (
	ErrCodeDivByZero ErrCode = iota
	ErrCodeInvalidInput
	ErrCodeUnknownOp
	ErrCodeInvalidTokenCount
)

// CalcError is the single error type used throughout the calculator.
type CalcError struct {
	Code    ErrCode
	Message string
	Wrapped error
}

func (e *CalcError) Error() string {
	return e.Message
}

func (e *CalcError) Unwrap() error {
	return e.Wrapped
}

// Sentinel errors for common calculator failures.
var (
	ErrDivByZero         = &CalcError{Code: ErrCodeDivByZero, Message: "Error: division by zero"}
	ErrInvalidInput      = &CalcError{Code: ErrCodeInvalidInput, Message: "Error: invalid numeric input"}
	ErrUnknownOp         = &CalcError{Code: ErrCodeUnknownOp, Message: "Error: unknown operator"}
	ErrInvalidTokenCount = &CalcError{Code: ErrCodeInvalidTokenCount, Message: "Error: usage: <number> <operator> <number>"}
)

// HandleError writes the error message to w.
func HandleError(err error, w io.Writer) {
	fmt.Fprintln(w, err.Error())
}
