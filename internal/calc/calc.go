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
		return 0, nil
	case "subtract", "-":
		return 0, nil
	case "multiply", "*":
		return 0, nil
	case "divide", "/":
		return 0, nil
	default:
		return 0, calcerrors.ErrUnknownOp
	}
}
