package calc_test

import (
	"errors"
	"testing"

	"github.com/repo/calculator/internal/calc"
	calcerrors "github.com/repo/calculator/internal/errors"
)

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
