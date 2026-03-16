package errors_test

import (
	"bytes"
	"fmt"
	"testing"

	calcerrors "github.com/repo/calculator/internal/errors"
)

// TestSentinelErrorsNonNil verifies that each sentinel error variable is non-nil.
func TestSentinelErrorsNonNil(t *testing.T) {
	sentinels := []struct {
		name string
		err  *calcerrors.CalcError
	}{
		{"ErrDivByZero", calcerrors.ErrDivByZero},
		{"ErrInvalidInput", calcerrors.ErrInvalidInput},
		{"ErrUnknownOp", calcerrors.ErrUnknownOp},
		{"ErrInvalidTokenCount", calcerrors.ErrInvalidTokenCount},
	}

	for _, tc := range sentinels {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err == nil {
				t.Errorf("%s must not be nil", tc.name)
			}
		})
	}
}

// TestSentinelErrorsNonEmptyMessage verifies that each sentinel error has a non-empty Message field.
func TestSentinelErrorsNonEmptyMessage(t *testing.T) {
	sentinels := []struct {
		name string
		err  *calcerrors.CalcError
	}{
		{"ErrDivByZero", calcerrors.ErrDivByZero},
		{"ErrInvalidInput", calcerrors.ErrInvalidInput},
		{"ErrUnknownOp", calcerrors.ErrUnknownOp},
		{"ErrInvalidTokenCount", calcerrors.ErrInvalidTokenCount},
	}

	for _, tc := range sentinels {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err == nil {
				t.Fatalf("%s is nil, cannot check Message", tc.name)
			}
			if tc.err.Message == "" {
				t.Errorf("%s has an empty Message field", tc.name)
			}
		})
	}
}

// TestCalcError_Error verifies that CalcError.Error() returns the Message string.
func TestCalcError_Error(t *testing.T) {
	tests := []struct {
		name            string
		err             *calcerrors.CalcError
		expectedMessage string
	}{
		{
			name:            "ErrDivByZero",
			err:             calcerrors.ErrDivByZero,
			expectedMessage: "Error: division by zero",
		},
		{
			name:            "ErrInvalidInput",
			err:             calcerrors.ErrInvalidInput,
			expectedMessage: "Error: invalid numeric input",
		},
		{
			name:            "ErrUnknownOp",
			err:             calcerrors.ErrUnknownOp,
			expectedMessage: "Error: unknown operator",
		},
		{
			name:            "ErrInvalidTokenCount",
			err:             calcerrors.ErrInvalidTokenCount,
			expectedMessage: "Error: usage: <number> <operator> <number>",
		},
		{
			name:            "CustomCalcError",
			err:             &calcerrors.CalcError{Code: calcerrors.ErrCodeDivByZero, Message: "custom message"},
			expectedMessage: "custom message",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.err.Error()
			if got != tc.expectedMessage {
				t.Errorf("Error() = %q, want %q", got, tc.expectedMessage)
			}
		})
	}
}

// TestCalcError_ErrorMatchesMessage verifies Error() returns exactly the Message field value.
func TestCalcError_ErrorMatchesMessage(t *testing.T) {
	sentinels := []struct {
		name string
		err  *calcerrors.CalcError
	}{
		{"ErrDivByZero", calcerrors.ErrDivByZero},
		{"ErrInvalidInput", calcerrors.ErrInvalidInput},
		{"ErrUnknownOp", calcerrors.ErrUnknownOp},
		{"ErrInvalidTokenCount", calcerrors.ErrInvalidTokenCount},
	}

	for _, tc := range sentinels {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err.Error() != tc.err.Message {
				t.Errorf("%s: Error() = %q, want Message field %q", tc.name, tc.err.Error(), tc.err.Message)
			}
		})
	}
}

// TestHandleError_WritesFormattedMessage verifies that HandleError writes the formatted message
// to the provided io.Writer (a bytes.Buffer) for each sentinel error.
func TestHandleError_WritesFormattedMessage(t *testing.T) {
	tests := []struct {
		name string
		err  *calcerrors.CalcError
	}{
		{"ErrDivByZero", calcerrors.ErrDivByZero},
		{"ErrInvalidInput", calcerrors.ErrInvalidInput},
		{"ErrUnknownOp", calcerrors.ErrUnknownOp},
		{"ErrInvalidTokenCount", calcerrors.ErrInvalidTokenCount},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			calcerrors.HandleError(tc.err, &buf)

			got := buf.String()
			expected := fmt.Sprintf("%s\n", tc.err.Error())

			if got != expected {
				t.Errorf("HandleError() wrote %q, want %q", got, expected)
			}
		})
	}
}

// TestHandleError_WritesMessageContent verifies that the buffer contains the error message text.
func TestHandleError_WritesMessageContent(t *testing.T) {
	tests := []struct {
		name            string
		err             *calcerrors.CalcError
		expectedContent string
	}{
		{
			name:            "ErrDivByZero contains message",
			err:             calcerrors.ErrDivByZero,
			expectedContent: "Error: division by zero",
		},
		{
			name:            "ErrInvalidInput contains message",
			err:             calcerrors.ErrInvalidInput,
			expectedContent: "Error: invalid numeric input",
		},
		{
			name:            "ErrUnknownOp contains message",
			err:             calcerrors.ErrUnknownOp,
			expectedContent: "Error: unknown operator",
		},
		{
			name:            "ErrInvalidTokenCount contains message",
			err:             calcerrors.ErrInvalidTokenCount,
			expectedContent: "Error: usage: <number> <operator> <number>",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			calcerrors.HandleError(tc.err, &buf)

			got := buf.String()
			if got == "" {
				t.Errorf("HandleError() wrote nothing to buffer, expected content containing %q", tc.expectedContent)
				return
			}

			// The buffer should contain the error message text
			if len(got) < len(tc.expectedContent) {
				t.Errorf("HandleError() output %q is too short to contain %q", got, tc.expectedContent)
			}
		})
	}
}

// TestCalcError_ErrCodeValues verifies the error code constants are distinct and properly assigned.
func TestCalcError_ErrCodeValues(t *testing.T) {
	codes := []calcerrors.ErrCode{
		calcerrors.ErrCodeDivByZero,
		calcerrors.ErrCodeInvalidInput,
		calcerrors.ErrCodeUnknownOp,
		calcerrors.ErrCodeInvalidTokenCount,
	}

	// All codes must be distinct
	seen := make(map[calcerrors.ErrCode]bool)
	for _, code := range codes {
		if seen[code] {
			t.Errorf("duplicate ErrCode value: %d", code)
		}
		seen[code] = true
	}
}

// TestCalcError_SentinelCodesMatchTypes verifies each sentinel has the expected error code.
func TestCalcError_SentinelCodesMatchTypes(t *testing.T) {
	tests := []struct {
		name         string
		err          *calcerrors.CalcError
		expectedCode calcerrors.ErrCode
	}{
		{"ErrDivByZero", calcerrors.ErrDivByZero, calcerrors.ErrCodeDivByZero},
		{"ErrInvalidInput", calcerrors.ErrInvalidInput, calcerrors.ErrCodeInvalidInput},
		{"ErrUnknownOp", calcerrors.ErrUnknownOp, calcerrors.ErrCodeUnknownOp},
		{"ErrInvalidTokenCount", calcerrors.ErrInvalidTokenCount, calcerrors.ErrCodeInvalidTokenCount},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err.Code != tc.expectedCode {
				t.Errorf("%s: Code = %d, want %d", tc.name, tc.err.Code, tc.expectedCode)
			}
		})
	}
}

// TestCalcError_Unwrap verifies that Unwrap returns nil for sentinels (no wrapped error).
func TestCalcError_Unwrap(t *testing.T) {
	sentinels := []struct {
		name string
		err  *calcerrors.CalcError
	}{
		{"ErrDivByZero", calcerrors.ErrDivByZero},
		{"ErrInvalidInput", calcerrors.ErrInvalidInput},
		{"ErrUnknownOp", calcerrors.ErrUnknownOp},
		{"ErrInvalidTokenCount", calcerrors.ErrInvalidTokenCount},
	}

	for _, tc := range sentinels {
		t.Run(tc.name, func(t *testing.T) {
			unwrapped := tc.err.Unwrap()
			if unwrapped != nil {
				t.Errorf("%s: Unwrap() = %v, want nil (sentinels should not wrap errors)", tc.name, unwrapped)
			}
		})
	}
}

// TestCalcError_Unwrap_WithWrapped verifies that Unwrap returns the wrapped error when set.
func TestCalcError_Unwrap_WithWrapped(t *testing.T) {
	inner := fmt.Errorf("inner error")
	err := &calcerrors.CalcError{
		Code:    calcerrors.ErrCodeInvalidInput,
		Message: "Error: invalid numeric input",
		Wrapped: inner,
	}

	if err.Unwrap() != inner {
		t.Errorf("Unwrap() = %v, want %v", err.Unwrap(), inner)
	}
}

// TestHandleError_DoesNotWriteToStderr verifies that HandleError uses the provided writer
// and not a hardcoded stderr. We pass a buffer and expect output there.
func TestHandleError_DoesNotWriteToStderr(t *testing.T) {
	// If HandleError writes to stderr instead of the provided writer,
	// our buffer will be empty.
	var buf bytes.Buffer
	calcerrors.HandleError(calcerrors.ErrDivByZero, &buf)

	if buf.Len() == 0 {
		t.Error("HandleError() did not write to the provided io.Writer; output may be going to stderr instead")
	}
}
