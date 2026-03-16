package calc

import (
	"math"

	calcerrors "github.com/repo/calculator/internal/errors"
)

// Calculate performs the arithmetic operation identified by op on operands a and b.
// Supported operators: "add", "+", "subtract", "-", "multiply", "*", "divide", "/"
// Returns (0, ErrUnknownOp) for unrecognised operators.
func Calculate(op string, a, b float64) (float64, error) {
	var result float64
	switch op {
	case "add", "+":
		result = a + b
	case "subtract", "-":
		result = a - b
	case "multiply", "*":
		result = a * b
	case "divide", "/":
		if b == 0 {
			return 0, calcerrors.ErrDivByZero
		}
		result = a / b
	default:
		return 0, calcerrors.ErrUnknownOp
	}
	if math.IsInf(result, 0) || math.IsNaN(result) {
		return 0, &calcerrors.CalcError{Code: calcerrors.ErrCodeInvalidInput, Message: "Error: result is not a finite number"}
	}
	return result, nil
}
