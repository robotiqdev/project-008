package calc_test

import (
	"errors"
	"testing"

	"github.com/repo/calculator/internal/calc"
	calcerrors "github.com/repo/calculator/internal/errors"
)

// TestCalculate_Add_TableDriven_WordForm tests the "add" operator with multiple cases.
func TestCalculate_Add_TableDriven_WordForm(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"positive integers", 1, 1, 2},
		{"both zero", 0, 0, 0},
		{"negative and positive", -1, 1, 0},
		{"floating point", 1.5, 2.5, 4.0},
		{"large values", 1e15, 2e15, 3e15},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Calculate("add", tc.a, tc.b)
			if err != nil {
				t.Fatalf("Calculate(\"add\", %v, %v): unexpected error: %v", tc.a, tc.b, err)
			}
			if result != tc.want {
				t.Errorf("Calculate(\"add\", %v, %v) = %v, want %v", tc.a, tc.b, result, tc.want)
			}
		})
	}
}

// TestCalculate_Add_TableDriven_SymbolicForm tests the "+" operator with multiple cases.
func TestCalculate_Add_TableDriven_SymbolicForm(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"positive integers", 1, 1, 2},
		{"both zero", 0, 0, 0},
		{"negative and positive", -1, 1, 0},
		{"floating point", 1.5, 2.5, 4.0},
		{"large values", 1e15, 2e15, 3e15},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Calculate("+", tc.a, tc.b)
			if err != nil {
				t.Fatalf("Calculate(\"+\", %v, %v): unexpected error: %v", tc.a, tc.b, err)
			}
			if result != tc.want {
				t.Errorf("Calculate(\"+\", %v, %v) = %v, want %v", tc.a, tc.b, result, tc.want)
			}
		})
	}
}

// TestCalculate_Subtract_TableDriven_WordForm tests the "subtract" operator with multiple cases.
func TestCalculate_Subtract_TableDriven_WordForm(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"positive result", 5, 3, 2},
		{"both zero", 0, 0, 0},
		{"negative result", 1, 2, -1},
		{"two negatives", -1, -1, 0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Calculate("subtract", tc.a, tc.b)
			if err != nil {
				t.Fatalf("Calculate(\"subtract\", %v, %v): unexpected error: %v", tc.a, tc.b, err)
			}
			if result != tc.want {
				t.Errorf("Calculate(\"subtract\", %v, %v) = %v, want %v", tc.a, tc.b, result, tc.want)
			}
		})
	}
}

// TestCalculate_Subtract_TableDriven_SymbolicForm tests the "-" operator with multiple cases.
func TestCalculate_Subtract_TableDriven_SymbolicForm(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"positive result", 5, 3, 2},
		{"both zero", 0, 0, 0},
		{"negative result", 1, 2, -1},
		{"two negatives", -1, -1, 0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Calculate("-", tc.a, tc.b)
			if err != nil {
				t.Fatalf("Calculate(\"-\", %v, %v): unexpected error: %v", tc.a, tc.b, err)
			}
			if result != tc.want {
				t.Errorf("Calculate(\"-\", %v, %v) = %v, want %v", tc.a, tc.b, result, tc.want)
			}
		})
	}
}

// TestCalculate_Add_WordForm verifies that the "add" operator returns the sum of two operands.
func TestCalculate_Add_WordForm(t *testing.T) {
	result, err := calc.Calculate("add", 1.0, 2.0)
	if err != nil {
		t.Fatalf("Calculate(\"add\", 1.0, 2.0): unexpected error: %v", err)
	}
	if result != 3.0 {
		t.Errorf("Calculate(\"add\", 1.0, 2.0) = %v, want 3.0", result)
	}
}

// TestCalculate_Add_SymbolicForm verifies that the "+" operator returns the sum of two operands.
func TestCalculate_Add_SymbolicForm(t *testing.T) {
	result, err := calc.Calculate("+", 5.0, 7.0)
	if err != nil {
		t.Fatalf("Calculate(\"+\", 5.0, 7.0): unexpected error: %v", err)
	}
	if result != 12.0 {
		t.Errorf("Calculate(\"+\", 5.0, 7.0) = %v, want 12.0", result)
	}
}

// TestCalculate_Subtract_WordForm verifies that the "subtract" operator returns the difference.
func TestCalculate_Subtract_WordForm(t *testing.T) {
	result, err := calc.Calculate("subtract", 10.0, 4.0)
	if err != nil {
		t.Fatalf("Calculate(\"subtract\", 10.0, 4.0): unexpected error: %v", err)
	}
	if result != 6.0 {
		t.Errorf("Calculate(\"subtract\", 10.0, 4.0) = %v, want 6.0", result)
	}
}

// TestCalculate_Subtract_SymbolicForm verifies that the "-" operator returns the difference.
func TestCalculate_Subtract_SymbolicForm(t *testing.T) {
	result, err := calc.Calculate("-", 9.0, 3.0)
	if err != nil {
		t.Fatalf("Calculate(\"-\", 9.0, 3.0): unexpected error: %v", err)
	}
	if result != 6.0 {
		t.Errorf("Calculate(\"-\", 9.0, 3.0) = %v, want 6.0", result)
	}
}

// TestCalculate_Multiply_WordForm verifies that the "multiply" operator returns the product.
func TestCalculate_Multiply_WordForm(t *testing.T) {
	result, err := calc.Calculate("multiply", 3.0, 4.0)
	if err != nil {
		t.Fatalf("Calculate(\"multiply\", 3.0, 4.0): unexpected error: %v", err)
	}
	if result != 12.0 {
		t.Errorf("Calculate(\"multiply\", 3.0, 4.0) = %v, want 12.0", result)
	}
}

// TestCalculate_Multiply_SymbolicForm verifies that the "*" operator returns the product.
func TestCalculate_Multiply_SymbolicForm(t *testing.T) {
	result, err := calc.Calculate("*", 6.0, 7.0)
	if err != nil {
		t.Fatalf("Calculate(\"*\", 6.0, 7.0): unexpected error: %v", err)
	}
	if result != 42.0 {
		t.Errorf("Calculate(\"*\", 6.0, 7.0) = %v, want 42.0", result)
	}
}

// TestCalculate_Divide_WordForm verifies that the "divide" operator returns the quotient.
func TestCalculate_Divide_WordForm(t *testing.T) {
	result, err := calc.Calculate("divide", 10.0, 2.0)
	if err != nil {
		t.Fatalf("Calculate(\"divide\", 10.0, 2.0): unexpected error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("Calculate(\"divide\", 10.0, 2.0) = %v, want 5.0", result)
	}
}

// TestCalculate_Divide_SymbolicForm verifies that the "/" operator returns the quotient.
func TestCalculate_Divide_SymbolicForm(t *testing.T) {
	result, err := calc.Calculate("/", 8.0, 4.0)
	if err != nil {
		t.Fatalf("Calculate(\"/\", 8.0, 4.0): unexpected error: %v", err)
	}
	if result != 2.0 {
		t.Errorf("Calculate(\"/\", 8.0, 4.0) = %v, want 2.0", result)
	}
}

// TestCalculate_UnknownOp_ReturnsErrUnknownOp verifies that an unrecognised operator
// causes Calculate to return (0, ErrUnknownOp).
func TestCalculate_UnknownOp_ReturnsErrUnknownOp(t *testing.T) {
	result, err := calc.Calculate("unknown", 1.0, 1.0)
	if result != 0 {
		t.Errorf("Calculate(\"unknown\", 1.0, 1.0) result = %v, want 0", result)
	}
	if err == nil {
		t.Fatal("Calculate(\"unknown\", 1.0, 1.0): expected ErrUnknownOp, got nil")
	}
	if !errors.Is(err, calcerrors.ErrUnknownOp) {
		t.Errorf("Calculate(\"unknown\", 1.0, 1.0): error = %v, want ErrUnknownOp", err)
	}
}

// TestCalculate_EmptyOp_ReturnsErrUnknownOp verifies that an empty operator string
// is treated as unknown.
func TestCalculate_EmptyOp_ReturnsErrUnknownOp(t *testing.T) {
	result, err := calc.Calculate("", 1.0, 1.0)
	if result != 0 {
		t.Errorf("Calculate(\"\", 1.0, 1.0) result = %v, want 0", result)
	}
	if !errors.Is(err, calcerrors.ErrUnknownOp) {
		t.Errorf("Calculate(\"\", 1.0, 1.0): error = %v, want ErrUnknownOp", err)
	}
}

// TestCalculate_CaseSensitive_UppercaseNotRecognised verifies that operator matching
// is case-sensitive (uppercase variants are not accepted).
func TestCalculate_CaseSensitive_UppercaseNotRecognised(t *testing.T) {
	for _, op := range []string{"ADD", "Add", "SUBTRACT", "MULTIPLY", "DIVIDE"} {
		op := op
		t.Run(op, func(t *testing.T) {
			result, err := calc.Calculate(op, 1.0, 1.0)
			if result != 0 {
				t.Errorf("Calculate(%q, 1.0, 1.0) result = %v, want 0", op, result)
			}
			if !errors.Is(err, calcerrors.ErrUnknownOp) {
				t.Errorf("Calculate(%q, 1.0, 1.0): error = %v, want ErrUnknownOp", op, err)
			}
		})
	}
}

// TestCalculate_KnownOps_ReturnNoError verifies that all recognised operator strings
// return a nil error (arithmetic correctness is checked by separate tests above).
func TestCalculate_KnownOps_ReturnNoError(t *testing.T) {
	knownOps := []string{"add", "+", "subtract", "-", "multiply", "*", "divide", "/"}
	for _, op := range knownOps {
		op := op
		t.Run(op, func(t *testing.T) {
			_, err := calc.Calculate(op, 2.0, 1.0)
			if err != nil {
				t.Errorf("Calculate(%q, 2.0, 1.0): unexpected error: %v", op, err)
			}
		})
	}
}

// TestCalculate_UnknownOp_ErrorMessage verifies the error message text for an unknown operator.
func TestCalculate_UnknownOp_ErrorMessage(t *testing.T) {
	_, err := calc.Calculate("modulo", 5.0, 2.0)
	if err == nil {
		t.Fatal("Calculate(\"modulo\", 5.0, 2.0): expected error, got nil")
	}
	want := "Error: unknown operator"
	if err.Error() != want {
		t.Errorf("error message = %q, want %q", err.Error(), want)
	}
}
