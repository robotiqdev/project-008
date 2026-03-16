package calc

import (
	calcerrors "github.com/repo/calculator/internal/errors"
)

// Calculate performs the arithmetic operation identified by op on operands a and b.
// Supported operators: "add", "+", "subtract", "-", "multiply", "*", "divide", "/"
// Returns (0, ErrUnknownOp) for unrecognised operators.
func Calculate(op string, a, b float64) (float64, error) {
	switch op {
	case "add", "+":
		return a + b, nil
	case "subtract", "-":
		return a - b, nil
	case "multiply", "*":
		return a * b, nil
	case "divide", "/":
		if b == 0 {
			return 0, calcerrors.ErrDivByZero
		}
		return a / b, nil
	default:
		return 0, calcerrors.ErrUnknownOp
	}
}
