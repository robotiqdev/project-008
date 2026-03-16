package calc

import (
	calcerrors "github.com/repo/calculator/internal/errors"
)

// Calculate applies op to a and b and returns the result.
func Calculate(op string, a, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, calcerrors.ErrDivByZero
		}
		return a / b, nil
	default:
		return 0, calcerrors.ErrUnknownOp
	}
}
