package calc_test

import (
	"errors"
	"math"
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

// TestCalculate_Multiply_TableDriven_WordForm tests the "multiply" operator with multiple cases.
func TestCalculate_Multiply_TableDriven_WordForm(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"positive integers", 3, 4, 12},
		{"multiply by zero", 0, 5, 0},
		{"negative and positive", -2, 3, -6},
		{"floating point", 1.5, 2, 3.0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Calculate("multiply", tc.a, tc.b)
			if err != nil {
				t.Fatalf("Calculate(\"multiply\", %v, %v): unexpected error: %v", tc.a, tc.b, err)
			}
			if result != tc.want {
				t.Errorf("Calculate(\"multiply\", %v, %v) = %v, want %v", tc.a, tc.b, result, tc.want)
			}
		})
	}
}

// TestCalculate_Multiply_TableDriven_SymbolicForm tests the "*" operator with multiple cases.
func TestCalculate_Multiply_TableDriven_SymbolicForm(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"positive integers", 3, 4, 12},
		{"multiply by zero", 0, 5, 0},
		{"negative and positive", -2, 3, -6},
		{"floating point", 1.5, 2, 3.0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Calculate("*", tc.a, tc.b)
			if err != nil {
				t.Fatalf("Calculate(\"*\", %v, %v): unexpected error: %v", tc.a, tc.b, err)
			}
			if result != tc.want {
				t.Errorf("Calculate(\"*\", %v, %v) = %v, want %v", tc.a, tc.b, result, tc.want)
			}
		})
	}
}

// TestCalculate_Divide_TableDriven_WordForm tests the "divide" operator with multiple cases.
func TestCalculate_Divide_TableDriven_WordForm(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		approx  bool
		epsilon float64
	}{
		{"exact integer result", 10, 2, 5, false, 0},
		{"fractional result", 7, 2, 3.5, false, 0},
		{"repeating decimal", 1, 3, 1.0 / 3.0, true, 1e-9},
		{"negative dividend", -6, 2, -3, false, 0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Calculate("divide", tc.a, tc.b)
			if err != nil {
				t.Fatalf("Calculate(\"divide\", %v, %v): unexpected error: %v", tc.a, tc.b, err)
			}
			if tc.approx {
				if math.Abs(result-tc.want) > tc.epsilon {
					t.Errorf("Calculate(\"divide\", %v, %v) = %v, want ~%v (within %v)", tc.a, tc.b, result, tc.want, tc.epsilon)
				}
			} else {
				if result != tc.want {
					t.Errorf("Calculate(\"divide\", %v, %v) = %v, want %v", tc.a, tc.b, result, tc.want)
				}
			}
		})
	}
}

// TestCalculate_Divide_TableDriven_SymbolicForm tests the "/" operator with multiple cases.
func TestCalculate_Divide_TableDriven_SymbolicForm(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		approx  bool
		epsilon float64
	}{
		{"exact integer result", 10, 2, 5, false, 0},
		{"fractional result", 7, 2, 3.5, false, 0},
		{"repeating decimal", 1, 3, 1.0 / 3.0, true, 1e-9},
		{"negative dividend", -6, 2, -3, false, 0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Calculate("/", tc.a, tc.b)
			if err != nil {
				t.Fatalf("Calculate(\"/\", %v, %v): unexpected error: %v", tc.a, tc.b, err)
			}
			if tc.approx {
				if math.Abs(result-tc.want) > tc.epsilon {
					t.Errorf("Calculate(\"/\", %v, %v) = %v, want ~%v (within %v)", tc.a, tc.b, result, tc.want, tc.epsilon)
				}
			} else {
				if result != tc.want {
					t.Errorf("Calculate(\"/\", %v, %v) = %v, want %v", tc.a, tc.b, result, tc.want)
				}
			}
		})
	}
}

// TestCalculate_Divide_ByZero_WordForm verifies that dividing by zero with "divide" returns ErrDivByZero.
func TestCalculate_Divide_ByZero_WordForm(t *testing.T) {
	result, err := calc.Calculate("divide", 5.0, 0)
	if result != 0 {
		t.Errorf("Calculate(\"divide\", 5.0, 0) result = %v, want 0", result)
	}
	if err == nil {
		t.Fatal("Calculate(\"divide\", 5.0, 0): expected ErrDivByZero, got nil")
	}
	if !errors.Is(err, calcerrors.ErrDivByZero) {
		t.Errorf("Calculate(\"divide\", 5.0, 0): error = %v, want ErrDivByZero", err)
	}
}

// TestCalculate_Divide_ByZero_SymbolicForm verifies that dividing by zero with "/" returns ErrDivByZero.
func TestCalculate_Divide_ByZero_SymbolicForm(t *testing.T) {
	result, err := calc.Calculate("/", 10.0, 0)
	if result != 0 {
		t.Errorf("Calculate(\"/\", 10.0, 0) result = %v, want 0", result)
	}
	if err == nil {
		t.Fatal("Calculate(\"/\", 10.0, 0): expected ErrDivByZero, got nil")
	}
	if !errors.Is(err, calcerrors.ErrDivByZero) {
		t.Errorf("Calculate(\"/\", 10.0, 0): error = %v, want ErrDivByZero", err)
	}
}

// TestCalculate_Divide_ByZero_ErrorMessage verifies the error message for division by zero.
func TestCalculate_Divide_ByZero_ErrorMessage(t *testing.T) {
	_, err := calc.Calculate("divide", 1.0, 0)
	if err == nil {
		t.Fatal("Calculate(\"divide\", 1.0, 0): expected error, got nil")
	}
	want := "Error: division by zero"
	if err.Error() != want {
		t.Errorf("error message = %q, want %q", err.Error(), want)
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

// --- Float64 edge-case tests (TASK-4278) ---
// Inf/NaN outputs are rejected: Calculate must return (0, error) whenever the
// arithmetic result would be +Inf, -Inf, or NaN, regardless of the operator.
// The error code must be ErrCodeInvalidInput and the message must be
// "Error: result is not a finite number".

// TestCalculate_Multiply_Overflow_MaxFloat64TimesTwo verifies that multiplying
// math.MaxFloat64 by 2 (which overflows to +Inf) returns (0, error) rather than +Inf.
func TestCalculate_Multiply_Overflow_MaxFloat64TimesTwo(t *testing.T) {
	result, err := calc.Calculate("multiply", math.MaxFloat64, 2)
	if result != 0 {
		t.Errorf("Calculate(\"multiply\", MaxFloat64, 2) result = %v, want 0", result)
	}
	if err == nil {
		t.Fatal("Calculate(\"multiply\", MaxFloat64, 2): expected error for overflow, got nil")
	}
	wantMsg := "Error: result is not a finite number"
	if err.Error() != wantMsg {
		t.Errorf("error message = %q, want %q", err.Error(), wantMsg)
	}
	var calcErr *calcerrors.CalcError
	if !errors.As(err, &calcErr) {
		t.Fatalf("error is not a *CalcError: %T", err)
	}
	if calcErr.Code != calcerrors.ErrCodeInvalidInput {
		t.Errorf("CalcError.Code = %v, want ErrCodeInvalidInput (%v)", calcErr.Code, calcerrors.ErrCodeInvalidInput)
	}
}

// TestCalculate_Multiply_Overflow_SymbolicForm verifies the same overflow behaviour
// through the "*" symbolic operator.
func TestCalculate_Multiply_Overflow_SymbolicForm(t *testing.T) {
	result, err := calc.Calculate("*", math.MaxFloat64, 2)
	if result != 0 {
		t.Errorf("Calculate(\"*\", MaxFloat64, 2) result = %v, want 0", result)
	}
	if err == nil {
		t.Fatal("Calculate(\"*\", MaxFloat64, 2): expected error for overflow, got nil")
	}
	wantMsg := "Error: result is not a finite number"
	if err.Error() != wantMsg {
		t.Errorf("error message = %q, want %q", err.Error(), wantMsg)
	}
}

// TestCalculate_Multiply_Overflow_NegativeInf verifies that multiplying
// -math.MaxFloat64 by 2 (which overflows to -Inf) is also rejected.
func TestCalculate_Multiply_Overflow_NegativeInf(t *testing.T) {
	result, err := calc.Calculate("multiply", -math.MaxFloat64, 2)
	if result != 0 {
		t.Errorf("Calculate(\"multiply\", -MaxFloat64, 2) result = %v, want 0", result)
	}
	if err == nil {
		t.Fatal("Calculate(\"multiply\", -MaxFloat64, 2): expected error for overflow, got nil")
	}
	wantMsg := "Error: result is not a finite number"
	if err.Error() != wantMsg {
		t.Errorf("error message = %q, want %q", err.Error(), wantMsg)
	}
}

// TestCalculate_Add_Overflow_MaxFloat64PlusMaxFloat64 verifies that adding two
// math.MaxFloat64 values (which overflows to +Inf) is rejected.
func TestCalculate_Add_Overflow_MaxFloat64PlusMaxFloat64(t *testing.T) {
	result, err := calc.Calculate("add", math.MaxFloat64, math.MaxFloat64)
	if result != 0 {
		t.Errorf("Calculate(\"add\", MaxFloat64, MaxFloat64) result = %v, want 0", result)
	}
	if err == nil {
		t.Fatal("Calculate(\"add\", MaxFloat64, MaxFloat64): expected error for overflow, got nil")
	}
	wantMsg := "Error: result is not a finite number"
	if err.Error() != wantMsg {
		t.Errorf("error message = %q, want %q", err.Error(), wantMsg)
	}
	var calcErr *calcerrors.CalcError
	if !errors.As(err, &calcErr) {
		t.Fatalf("error is not a *CalcError: %T", err)
	}
	if calcErr.Code != calcerrors.ErrCodeInvalidInput {
		t.Errorf("CalcError.Code = %v, want ErrCodeInvalidInput (%v)", calcErr.Code, calcerrors.ErrCodeInvalidInput)
	}
}

// TestCalculate_Multiply_ZeroTimesZero_Valid verifies that 0 * 0 = 0 is accepted as valid.
func TestCalculate_Multiply_ZeroTimesZero_Valid(t *testing.T) {
	result, err := calc.Calculate("multiply", 0, 0)
	if err != nil {
		t.Fatalf("Calculate(\"multiply\", 0, 0): unexpected error: %v", err)
	}
	if result != 0 {
		t.Errorf("Calculate(\"multiply\", 0, 0) = %v, want 0", result)
	}
}

// TestCalculate_Multiply_VerySmallNumbers_Valid verifies that operations on very
// small finite numbers near zero do not cause errors (underflow to 0 is acceptable).
func TestCalculate_Multiply_VerySmallNumbers_Valid(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
	}{
		{"SmallestNonzero * SmallestNonzero", math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64},
		{"SmallestNonzero * 1", math.SmallestNonzeroFloat64, 1},
		{"1e-300 * 1e-300", 1e-300, 1e-300},
		{"1e-200 * 1e-200", 1e-200, 1e-200},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Underflow to 0 produces a finite result (0), which is valid.
			result, err := calc.Calculate("multiply", tc.a, tc.b)
			if err != nil {
				t.Fatalf("Calculate(\"multiply\", %v, %v): unexpected error: %v", tc.a, tc.b, err)
			}
			if math.IsInf(result, 0) || math.IsNaN(result) {
				t.Errorf("Calculate(\"multiply\", %v, %v) = %v, want finite result", tc.a, tc.b, result)
			}
		})
	}
}

// TestCalculate_Multiply_VerySmallNumbers_ResultIsFinite verifies that multiplying
// very small numbers with moderate values stays finite.
func TestCalculate_Multiply_VerySmallNumbers_ResultIsFinite(t *testing.T) {
	result, err := calc.Calculate("multiply", 1e-150, 1e-150)
	if err != nil {
		t.Fatalf("Calculate(\"multiply\", 1e-150, 1e-150): unexpected error: %v", err)
	}
	if math.IsInf(result, 0) || math.IsNaN(result) {
		t.Errorf("Calculate(\"multiply\", 1e-150, 1e-150) = %v, want finite result", result)
	}
}

// TestCalculate_Overflow_ResultIsNotReturned verifies that when overflow produces +Inf,
// the returned result value is 0 (not +Inf or any non-zero value).
func TestCalculate_Overflow_ResultIsNotReturned(t *testing.T) {
	overflowCases := []struct {
		op string
		a  float64
		b  float64
	}{
		{"multiply", math.MaxFloat64, 2},
		{"*", math.MaxFloat64, 3},
		{"add", math.MaxFloat64, math.MaxFloat64},
		{"+", math.MaxFloat64, math.MaxFloat64},
	}
	for _, tc := range overflowCases {
		tc := tc
		t.Run(tc.op, func(t *testing.T) {
			result, err := calc.Calculate(tc.op, tc.a, tc.b)
			if err == nil {
				t.Errorf("Calculate(%q, %v, %v): expected error, got nil", tc.op, tc.a, tc.b)
			}
			if result != 0 {
				t.Errorf("Calculate(%q, %v, %v) result = %v, want 0 on overflow", tc.op, tc.a, tc.b, result)
			}
			if math.IsInf(result, 0) || math.IsNaN(result) {
				t.Errorf("Calculate(%q, %v, %v) = %v, Inf/NaN must never be returned", tc.op, tc.a, tc.b, result)
			}
		})
	}
}
