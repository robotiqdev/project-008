package calc

// Calculate performs the arithmetic operation identified by op on operands a and b.
// Supported ops: "add", "+", "subtract", "-", "multiply", "*", "divide", "/".
// Returns the result and nil on success, or 0 and a non-nil error on failure.
func Calculate(op string, a, b float64) (float64, error) {
	switch op {
	case "add", "+":
		return a + b, nil
	case "subtract", "-":
		return a - b, nil
	case "multiply", "*":
		return a * b, nil
	case "divide", "/":
		return divide(a, b)
	default:
		return 0, nil
	}
}

func divide(a, b float64) (float64, error) {
	if err := validateDivisor(b); err != nil {
		return 0, err
	}
	return a / b, nil
}
