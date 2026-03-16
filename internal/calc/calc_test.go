package calc_test

import (
	"errors"
	"math"
	"testing"

	"github.com/repo/calculator/internal/calc"
	calcerrors "github.com/repo/calculator/internal/errors"
)

// --- Table-driven tests for all operators and edge cases ---

type calcTestCase struct {
	name        string
	op          string
	a, b        float64
	wantResult  float64
	wantErr     error
	wantErrCode calcerrors.ErrCode
	wantInf     bool // true if result should be +/- Inf
}

func TestCalculate_TableDriven(t *testing.T) {
	tests := []calcTestCase{
		// --- "add" operator ---
		{name: "add_positive_operands", op: "add", a: 3, b: 4, wantResult: 7},
		{name: "add_zero_operands", op: "add", a: 0, b: 0, wantResult: 0},
		{name: "add_zero_left_operand", op: "add", a: 0, b: 5, wantResult: 5},
		{name: "add_zero_right_operand", op: "add", a: 5, b: 0, wantResult: 5},
		{name: "add_negative_operands", op: "add", a: -3, b: -4, wantResult: -7},
		{name: "add_negative_and_positive", op: "add", a: -3, b: 4, wantResult: 1},
		{name: "add_floats", op: "add", a: 1.5, b: 2.5, wantResult: 4.0},

		// --- "+" operator ---
		{name: "plus_positive_operands", op: "+", a: 10, b: 20, wantResult: 30},
		{name: "plus_zero_operands", op: "+", a: 0, b: 0, wantResult: 0},
		{name: "plus_zero_left_operand", op: "+", a: 0, b: 7, wantResult: 7},
		{name: "plus_zero_right_operand", op: "+", a: 7, b: 0, wantResult: 7},
		{name: "plus_negative_operands", op: "+", a: -5, b: -5, wantResult: -10},
		{name: "plus_negative_and_positive", op: "+", a: -10, b: 3, wantResult: -7},

		// --- "subtract" operator ---
		{name: "subtract_positive_operands", op: "subtract", a: 10, b: 4, wantResult: 6},
		{name: "subtract_zero_operands", op: "subtract", a: 0, b: 0, wantResult: 0},
		{name: "subtract_zero_left_operand", op: "subtract", a: 0, b: 5, wantResult: -5},
		{name: "subtract_zero_right_operand", op: "subtract", a: 5, b: 0, wantResult: 5},
		{name: "subtract_negative_operands", op: "subtract", a: -3, b: -4, wantResult: 1},
		{name: "subtract_negative_from_positive", op: "subtract", a: 5, b: -3, wantResult: 8},
		{name: "subtract_floats", op: "subtract", a: 3.5, b: 1.5, wantResult: 2.0},

		// --- "-" operator ---
		{name: "minus_positive_operands", op: "-", a: 20, b: 8, wantResult: 12},
		{name: "minus_zero_operands", op: "-", a: 0, b: 0, wantResult: 0},
		{name: "minus_zero_left_operand", op: "-", a: 0, b: 3, wantResult: -3},
		{name: "minus_zero_right_operand", op: "-", a: 3, b: 0, wantResult: 3},
		{name: "minus_negative_operands", op: "-", a: -6, b: -2, wantResult: -4},

		// --- "multiply" operator ---
		{name: "multiply_positive_operands", op: "multiply", a: 3, b: 4, wantResult: 12},
		{name: "multiply_zero_left_operand", op: "multiply", a: 0, b: 5, wantResult: 0},
		{name: "multiply_zero_right_operand", op: "multiply", a: 5, b: 0, wantResult: 0},
		{name: "multiply_zero_operands", op: "multiply", a: 0, b: 0, wantResult: 0},
		{name: "multiply_negative_operands", op: "multiply", a: -3, b: -4, wantResult: 12},
		{name: "multiply_negative_and_positive", op: "multiply", a: -3, b: 4, wantResult: -12},
		{name: "multiply_floats", op: "multiply", a: 2.5, b: 4.0, wantResult: 10.0},

		// --- "*" operator ---
		{name: "star_positive_operands", op: "*", a: 6, b: 7, wantResult: 42},
		{name: "star_zero_left_operand", op: "*", a: 0, b: 9, wantResult: 0},
		{name: "star_zero_right_operand", op: "*", a: 9, b: 0, wantResult: 0},
		{name: "star_zero_operands", op: "*", a: 0, b: 0, wantResult: 0},
		{name: "star_negative_operands", op: "*", a: -5, b: -5, wantResult: 25},
		{name: "star_negative_and_positive", op: "*", a: -5, b: 5, wantResult: -25},

		// --- "divide" operator (non-zero divisor) ---
		{name: "divide_positive_operands", op: "divide", a: 10, b: 2, wantResult: 5},
		{name: "divide_zero_dividend", op: "divide", a: 0, b: 5, wantResult: 0},
		{name: "divide_negative_operands", op: "divide", a: -10, b: -2, wantResult: 5},
		{name: "divide_negative_dividend", op: "divide", a: -10, b: 2, wantResult: -5},
		{name: "divide_negative_divisor", op: "divide", a: 10, b: -2, wantResult: -5},
		{name: "divide_floats", op: "divide", a: 7.5, b: 2.5, wantResult: 3.0},
		{name: "divide_near_zero_divisor_no_error", op: "divide", a: 5, b: 0.001, wantResult: 5000},

		// --- "/" operator (non-zero divisor) ---
		{name: "slash_positive_operands", op: "/", a: 15, b: 3, wantResult: 5},
		{name: "slash_zero_dividend", op: "/", a: 0, b: 4, wantResult: 0},
		{name: "slash_negative_operands", op: "/", a: -12, b: -3, wantResult: 4},
		{name: "slash_negative_dividend", op: "/", a: -12, b: 3, wantResult: -4},
		{name: "slash_negative_divisor", op: "/", a: 12, b: -3, wantResult: -4},

		// --- div-by-zero errors ---
		{name: "divide_by_zero_using_divide_op", op: "divide", a: 5, b: 0, wantErr: calcerrors.ErrDivByZero},
		{name: "divide_by_zero_using_slash_op", op: "/", a: 5, b: 0, wantErr: calcerrors.ErrDivByZero},
		{name: "divide_zero_by_zero_using_divide_op", op: "divide", a: 0, b: 0, wantErr: calcerrors.ErrDivByZero},
		{name: "divide_zero_by_zero_using_slash_op", op: "/", a: 0, b: 0, wantErr: calcerrors.ErrDivByZero},
		{name: "divide_negative_by_zero_using_divide_op", op: "divide", a: -5, b: 0, wantErr: calcerrors.ErrDivByZero},
		{name: "divide_negative_by_zero_using_slash_op", op: "/", a: -5, b: 0, wantErr: calcerrors.ErrDivByZero},

		// --- unknown operator ---
		{name: "unknown_op_returns_ErrUnknownOp", op: "modulo", a: 5, b: 3, wantErr: calcerrors.ErrUnknownOp},
		{name: "empty_op_returns_ErrUnknownOp", op: "", a: 5, b: 3, wantErr: calcerrors.ErrUnknownOp},
		{name: "uppercase_ADD_returns_ErrUnknownOp", op: "ADD", a: 5, b: 3, wantErr: calcerrors.ErrUnknownOp},
		{name: "percent_op_returns_ErrUnknownOp", op: "%", a: 5, b: 3, wantErr: calcerrors.ErrUnknownOp},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Calculate(tc.op, tc.a, tc.b)

			if tc.wantErr != nil {
				// Expect an error
				if err == nil {
					t.Fatalf("Calculate(%q, %v, %v) expected error %v, got nil", tc.op, tc.a, tc.b, tc.wantErr)
				}
				if !errors.Is(err, tc.wantErr) {
					t.Errorf("Calculate(%q, %v, %v) error = %v, want errors.Is to match %v", tc.op, tc.a, tc.b, err, tc.wantErr)
				}
				if result != 0 {
					t.Errorf("Calculate(%q, %v, %v) result = %v on error, want 0", tc.op, tc.a, tc.b, result)
				}
				return
			}

			// Expect no error
			if err != nil {
				t.Fatalf("Calculate(%q, %v, %v) unexpected error: %v", tc.op, tc.a, tc.b, err)
			}

			if tc.wantInf {
				if !math.IsInf(result, 0) {
					t.Errorf("Calculate(%q, %v, %v) = %v, want Inf", tc.op, tc.a, tc.b, result)
				}
			} else {
				if result != tc.wantResult {
					t.Errorf("Calculate(%q, %v, %v) = %v, want %v", tc.op, tc.a, tc.b, result, tc.wantResult)
				}
			}
		})
	}
}

// --- Explicit tests retained from previous iteration ---

// TestCalculate_Divide_ByZero_ReturnsErrDivByZero verifies that Calculate("divide", 5.0, 0.0)
// returns (0, err) where errors.Is(err, ErrDivByZero).
func TestCalculate_Divide_ByZero_ReturnsErrDivByZero(t *testing.T) {
	result, err := calc.Calculate("divide", 5.0, 0.0)

	if err == nil {
		t.Fatal("Calculate(\"divide\", 5.0, 0.0) expected an error, got nil")
	}
	if !errors.Is(err, calcerrors.ErrDivByZero) {
		t.Errorf("Calculate(\"divide\", 5.0, 0.0) error = %v, want errors.Is to match ErrDivByZero", err)
	}
	if result != 0 {
		t.Errorf("Calculate(\"divide\", 5.0, 0.0) result = %v, want 0", result)
	}
}

// TestCalculate_SlashOp_ByZero_ReturnsErrDivByZero verifies that Calculate("/", 5.0, 0.0)
// returns (0, err) where errors.Is(err, ErrDivByZero).
func TestCalculate_SlashOp_ByZero_ReturnsErrDivByZero(t *testing.T) {
	result, err := calc.Calculate("/", 5.0, 0.0)

	if err == nil {
		t.Fatal("Calculate(\"/\", 5.0, 0.0) expected an error, got nil")
	}
	if !errors.Is(err, calcerrors.ErrDivByZero) {
		t.Errorf("Calculate(\"/\", 5.0, 0.0) error = %v, want errors.Is to match ErrDivByZero", err)
	}
	if result != 0 {
		t.Errorf("Calculate(\"/\", 5.0, 0.0) result = %v, want 0", result)
	}
}

// TestCalculate_Divide_NearZeroDivisor_NoError verifies that Calculate("divide", 5.0, 0.001)
// succeeds — 0.001 is not exactly zero and must not trigger ErrDivByZero.
func TestCalculate_Divide_NearZeroDivisor_NoError(t *testing.T) {
	_, err := calc.Calculate("divide", 5.0, 0.001)

	if err != nil {
		t.Errorf("Calculate(\"divide\", 5.0, 0.001) expected no error, got %v", err)
	}
}

// TestCalculate_Divide_ZeroDividedByZero_ReturnsErrDivByZero verifies that
// Calculate("divide", 0.0, 0.0) returns (0, err) where errors.Is(err, ErrDivByZero).
func TestCalculate_Divide_ZeroDividedByZero_ReturnsErrDivByZero(t *testing.T) {
	result, err := calc.Calculate("divide", 0.0, 0.0)

	if err == nil {
		t.Fatal("Calculate(\"divide\", 0.0, 0.0) expected an error, got nil")
	}
	if !errors.Is(err, calcerrors.ErrDivByZero) {
		t.Errorf("Calculate(\"divide\", 0.0, 0.0) error = %v, want errors.Is to match ErrDivByZero", err)
	}
	if result != 0 {
		t.Errorf("Calculate(\"divide\", 0.0, 0.0) result = %v, want 0", result)
	}
}

// --- Explicit unknown-operator tests ---

// TestCalculate_UnknownOp_ReturnsErrUnknownOp verifies that an unrecognised operator
// returns (0, ErrUnknownOp).
func TestCalculate_UnknownOp_ReturnsErrUnknownOp(t *testing.T) {
	result, err := calc.Calculate("modulo", 10, 3)

	if err == nil {
		t.Fatal("Calculate(\"modulo\", 10, 3) expected ErrUnknownOp, got nil")
	}
	if !errors.Is(err, calcerrors.ErrUnknownOp) {
		t.Errorf("Calculate(\"modulo\", 10, 3) error = %v, want errors.Is to match ErrUnknownOp", err)
	}
	if result != 0 {
		t.Errorf("Calculate(\"modulo\", 10, 3) result = %v on error, want 0", result)
	}
}

// TestCalculate_EmptyOp_ReturnsErrUnknownOp verifies that an empty operator string
// returns (0, ErrUnknownOp).
func TestCalculate_EmptyOp_ReturnsErrUnknownOp(t *testing.T) {
	result, err := calc.Calculate("", 1, 2)

	if err == nil {
		t.Fatal("Calculate(\"\", 1, 2) expected ErrUnknownOp, got nil")
	}
	if !errors.Is(err, calcerrors.ErrUnknownOp) {
		t.Errorf("Calculate(\"\", 1, 2) error = %v, want errors.Is to match ErrUnknownOp", err)
	}
	if result != 0 {
		t.Errorf("Calculate(\"\", 1, 2) result = %v on error, want 0", result)
	}
}

// --- Float64 edge-case tests ---

// TestCalculate_Add_ProducesInf_WhenResultOverflows verifies that adding two very large
// float64 values produces +Inf (Go's IEEE-754 behaviour — no error expected).
func TestCalculate_Add_ProducesInf_WhenResultOverflows(t *testing.T) {
	result, err := calc.Calculate("add", math.MaxFloat64, math.MaxFloat64)

	if err != nil {
		t.Fatalf("Calculate(\"add\", MaxFloat64, MaxFloat64) unexpected error: %v", err)
	}
	if !math.IsInf(result, 1) {
		t.Errorf("Calculate(\"add\", MaxFloat64, MaxFloat64) = %v, want +Inf", result)
	}
}

// TestCalculate_Multiply_ProducesInf_WhenResultOverflows verifies that multiplying two
// very large float64 values produces +Inf.
func TestCalculate_Multiply_ProducesInf_WhenResultOverflows(t *testing.T) {
	result, err := calc.Calculate("multiply", math.MaxFloat64, 2)

	if err != nil {
		t.Fatalf("Calculate(\"multiply\", MaxFloat64, 2) unexpected error: %v", err)
	}
	if !math.IsInf(result, 1) {
		t.Errorf("Calculate(\"multiply\", MaxFloat64, 2) = %v, want +Inf", result)
	}
}
