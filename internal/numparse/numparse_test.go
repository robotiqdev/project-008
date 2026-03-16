package numparse_test

import (
	"errors"
	"math"
	"strings"
	"testing"

	calcerrors "github.com/robotiqdev/project-008/internal/errors"
	"github.com/robotiqdev/project-008/internal/numparse"
)

const epsilon = 1e-9

func TestParseNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		// Valid integer strings
		{name: "zero", input: "0", want: 0, wantErr: false},
		{name: "positive integer", input: "42", want: 42, wantErr: false},
		{name: "negative integer", input: "-7", want: -7, wantErr: false},
		{name: "large integer", input: "100", want: 100, wantErr: false},

		// Valid decimal strings
		{name: "pi decimal", input: "3.14", want: 3.14, wantErr: false},
		{name: "negative decimal", input: "-0.5", want: -0.5, wantErr: false},
		{name: "one point zero", input: "1.0", want: 1.0, wantErr: false},

		// Invalid / non-numeric strings
		{name: "alphabetic string", input: "abc", want: 0, wantErr: true},
		{name: "empty string", input: "", want: 0, wantErr: true},
		{name: "word number", input: "one", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := numparse.ParseNumber(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseNumber(%q) expected an error, got nil (value=%v)", tt.input, got)
				}
				if got != 0 {
					t.Errorf("ParseNumber(%q) on error expected return value 0, got %v", tt.input, got)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseNumber(%q) unexpected error: %v", tt.input, err)
				return
			}

			if math.Abs(got-tt.want) > epsilon {
				t.Errorf("ParseNumber(%q) = %v, want %v (diff=%v)", tt.input, got, tt.want, math.Abs(got-tt.want))
			}
		})
	}
}

// TestParseNumber_ReturnsErrInvalidInput verifies that parsing a non-numeric string
// returns an error satisfying errors.Is(err, calcerrors.ErrInvalidInput).
func TestParseNumber_ReturnsErrInvalidInput(t *testing.T) {
	_, err := numparse.ParseNumber("abc")
	if err == nil {
		t.Fatal("ParseNumber(\"abc\") expected an error, got nil")
	}
	if !errors.Is(err, calcerrors.ErrInvalidInput) {
		t.Errorf("ParseNumber(\"abc\") error should satisfy errors.Is(err, ErrInvalidInput), got: %v (type: %T)", err, err)
	}
}

// TestParseNumber_ErrorMessageContainsInput verifies that the error message
// contains the offending input string for better UX.
func TestParseNumber_ErrorMessageContainsInput(t *testing.T) {
	input := "abc"
	_, err := numparse.ParseNumber(input)
	if err == nil {
		t.Fatal("ParseNumber(\"abc\") expected an error, got nil")
	}
	if !strings.Contains(err.Error(), input) {
		t.Errorf("ParseNumber(%q) error message %q should contain the input %q", input, err.Error(), input)
	}
}

// TestParseNumber_EmptyString_ErrorMessageContainsInput verifies that the error
// message for an empty string input contains the offending input for better UX.
func TestParseNumber_EmptyString_ErrorMessageContainsInput(t *testing.T) {
	input := ""
	_, err := numparse.ParseNumber(input)
	if err == nil {
		t.Fatal("ParseNumber(\"\") expected an error, got nil")
	}
	// The error message should include the quoted empty string representation
	if !strings.Contains(err.Error(), `""`) {
		t.Errorf("ParseNumber(%q) error message %q should contain quoted empty string", input, err.Error())
	}
}

// TestValidate_EmptyString_ReturnsErrInvalidInput verifies that Validate("")
// returns an error satisfying errors.Is(err, ErrInvalidInput).
func TestValidate_EmptyString_ReturnsErrInvalidInput(t *testing.T) {
	err := numparse.Validate("")
	if err == nil {
		t.Fatal("Validate(\"\") expected an error, got nil")
	}
	if !errors.Is(err, calcerrors.ErrInvalidInput) {
		t.Errorf("Validate(\"\") error should satisfy errors.Is(err, ErrInvalidInput), got: %v (type: %T)", err, err)
	}
}

// TestValidate_ValidDecimal_ReturnsNil verifies that Validate("3.14") returns nil.
func TestValidate_ValidDecimal_ReturnsNil(t *testing.T) {
	err := numparse.Validate("3.14")
	if err != nil {
		t.Errorf("Validate(\"3.14\") expected nil error, got: %v", err)
	}
}

// TestValidate_InvalidInput_ReturnsErrInvalidInput verifies that Validate with
// a non-numeric string returns ErrInvalidInput.
func TestValidate_InvalidInput_ReturnsErrInvalidInput(t *testing.T) {
	err := numparse.Validate("not-a-number")
	if err == nil {
		t.Fatal("Validate(\"not-a-number\") expected an error, got nil")
	}
	if !errors.Is(err, calcerrors.ErrInvalidInput) {
		t.Errorf("Validate(\"not-a-number\") error should satisfy errors.Is(err, ErrInvalidInput), got: %v (type: %T)", err, err)
	}
}

// TestValidate_ValidInteger_ReturnsNil verifies that Validate with a valid integer
// string returns nil.
func TestValidate_ValidInteger_ReturnsNil(t *testing.T) {
	err := numparse.Validate("42")
	if err != nil {
		t.Errorf("Validate(\"42\") expected nil error, got: %v", err)
	}
}
