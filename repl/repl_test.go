package repl_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/repo/calculator/repl"
)

// TestRun_DivByZero_WritesErrorToErrOut verifies that when the input "10 / 0\n" is
// provided, the errOut buffer receives the division-by-zero error message.
func TestRun_DivByZero_WritesErrorToErrOut(t *testing.T) {
	in := strings.NewReader("10 / 0\n")
	var out, errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	got := errOut.String()
	want := "Error: division by zero\n"
	if got != want {
		t.Errorf("errOut = %q, want %q", got, want)
	}
}

// TestRun_DivByZero_DoesNotWriteResultToOut verifies that a division-by-zero error
// does not produce any output on the normal out buffer.
func TestRun_DivByZero_DoesNotWriteResultToOut(t *testing.T) {
	in := strings.NewReader("10 / 0\n")
	var out, errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	if out.Len() != 0 {
		t.Errorf("out = %q, want empty (no result for error input)", out.String())
	}
}

// TestRun_AfterDivByZero_ContinuesWithValidInput verifies that after a division-by-zero
// error, the REPL continues processing subsequent lines. Input "6 / 2\n" following
// "10 / 0\n" must produce "3" on out.
func TestRun_AfterDivByZero_ContinuesWithValidInput(t *testing.T) {
	in := strings.NewReader("10 / 0\n6 / 2\n")
	var out, errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	gotOut := out.String()
	if !strings.Contains(gotOut, "3") {
		t.Errorf("out = %q, want it to contain \"3\" (result of 6 / 2)", gotOut)
	}
}

// TestRun_BothBuffers_DivByZeroThenValid verifies the combined behavior:
// "10 / 0\n6 / 2\n" — errOut receives the div-by-zero message AND out receives "3".
func TestRun_BothBuffers_DivByZeroThenValid(t *testing.T) {
	in := strings.NewReader("10 / 0\n6 / 2\n")
	var out, errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	gotErr := errOut.String()
	wantErr := "Error: division by zero\n"
	if gotErr != wantErr {
		t.Errorf("errOut = %q, want %q", gotErr, wantErr)
	}

	gotOut := out.String()
	if !strings.Contains(gotOut, "3") {
		t.Errorf("out = %q, want it to contain \"3\" (result of 6 / 2)", gotOut)
	}
}

// TestRun_ValidDivision_WritesResultToOut verifies that a valid division expression
// writes its result to out and nothing to errOut.
func TestRun_ValidDivision_WritesResultToOut(t *testing.T) {
	in := strings.NewReader("6 / 2\n")
	var out, errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	gotOut := out.String()
	if !strings.Contains(gotOut, "3") {
		t.Errorf("out = %q, want it to contain \"3\"", gotOut)
	}

	if errOut.Len() != 0 {
		t.Errorf("errOut = %q, want empty for successful calculation", errOut.String())
	}
}

// TestRun_DivByZero_ErrOutExactMessage verifies the exact format of the error message
// written to errOut for division by zero.
func TestRun_DivByZero_ErrOutExactMessage(t *testing.T) {
	in := strings.NewReader("10 / 0\n")
	var out, errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	got := errOut.String()
	want := "Error: division by zero\n"
	if got != want {
		t.Errorf("errOut = %q, want exactly %q", got, want)
	}
}

// TestRun_MultipleDivByZero_AllErrorsGoToErrOut verifies that multiple consecutive
// division-by-zero errors all route to errOut, not out.
func TestRun_MultipleDivByZero_AllErrorsGoToErrOut(t *testing.T) {
	in := strings.NewReader("10 / 0\n5 / 0\n")
	var out, errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	if out.Len() != 0 {
		t.Errorf("out = %q, want empty (no valid results)", out.String())
	}

	gotErr := errOut.String()
	wantErr := "Error: division by zero\nError: division by zero\n"
	if gotErr != wantErr {
		t.Errorf("errOut = %q, want %q", gotErr, wantErr)
	}
}

// TestRun_ErrorThenMultipleValid_ContinuesCorrectly verifies that after a div-by-zero
// error, the REPL continues processing multiple subsequent valid lines.
func TestRun_ErrorThenMultipleValid_ContinuesCorrectly(t *testing.T) {
	in := strings.NewReader("10 / 0\n6 / 2\n4 / 2\n")
	var out, errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	gotErr := errOut.String()
	wantErr := "Error: division by zero\n"
	if gotErr != wantErr {
		t.Errorf("errOut = %q, want %q", gotErr, wantErr)
	}

	gotOut := out.String()
	if !strings.Contains(gotOut, "3") {
		t.Errorf("out = %q, want it to contain \"3\" (result of 6 / 2)", gotOut)
	}
	if !strings.Contains(gotOut, "2") {
		t.Errorf("out = %q, want it to contain \"2\" (result of 4 / 2)", gotOut)
	}
}
