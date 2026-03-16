package repl_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/repo/calculator/repl"
)

// TestRun_AddIntegers verifies that '3 + 4\n' produces '7\n' in the output buffer.
func TestRun_AddIntegers(t *testing.T) {
	in := strings.NewReader("3 + 4\nexit\n")
	var out bytes.Buffer
	var errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	got := out.String()
	want := "7\n"
	if got != want {
		t.Errorf("Run(\"3 + 4\\nexit\\n\") out = %q, want %q", got, want)
	}
}

// TestRun_DivideRepeatingDecimal verifies that '1 / 3\n' produces '0.3333333333333333\n'.
func TestRun_DivideRepeatingDecimal(t *testing.T) {
	in := strings.NewReader("1 / 3\nexit\n")
	var out bytes.Buffer
	var errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	got := out.String()
	want := "0.3333333333333333\n"
	if got != want {
		t.Errorf("Run(\"1 / 3\\nexit\\n\") out = %q, want %q", got, want)
	}
}

// TestRun_MultiplyFloatByInteger verifies that '2.5 * 2\n' produces '5\n'.
// strconv.FormatFloat with 'f' format and -1 precision omits trailing zeros,
// so 5.0 is formatted as "5", not "5.0" or "5.000000".
func TestRun_MultiplyFloatByInteger(t *testing.T) {
	in := strings.NewReader("2.5 * 2\nexit\n")
	var out bytes.Buffer
	var errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	got := out.String()
	want := "5\n"
	if got != want {
		t.Errorf("Run(\"2.5 * 2\\nexit\\n\") out = %q, want %q", got, want)
	}
}

// TestRun_SubtractIntegers verifies that '10 - 3\n' produces '7\n' in the output buffer.
func TestRun_SubtractIntegers(t *testing.T) {
	in := strings.NewReader("10 - 3\nexit\n")
	var out bytes.Buffer
	var errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	got := out.String()
	want := "7\n"
	if got != want {
		t.Errorf("Run(\"10 - 3\\nexit\\n\") out = %q, want %q", got, want)
	}
}

// TestRun_ExitWritesNothingToOut verifies that 'exit\n' causes Run to return
// without writing anything to the out buffer.
func TestRun_ExitWritesNothingToOut(t *testing.T) {
	in := strings.NewReader("exit\n")
	var out bytes.Buffer
	var errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	if out.Len() != 0 {
		t.Errorf("Run(\"exit\\n\") out = %q, want empty string", out.String())
	}
}

// TestRun_MultipleExpressions verifies that multiple expressions are evaluated
// in sequence, each producing its own line of output.
func TestRun_MultipleExpressions(t *testing.T) {
	in := strings.NewReader("3 + 4\n10 - 3\nexit\n")
	var out bytes.Buffer
	var errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	got := out.String()
	want := "7\n7\n"
	if got != want {
		t.Errorf("Run multiple expressions: out = %q, want %q", got, want)
	}
}

// TestRun_IntegerDivisionExact verifies that exact integer division produces
// an integer-formatted result (no trailing decimal point or zeros).
func TestRun_IntegerDivisionExact(t *testing.T) {
	in := strings.NewReader("10 / 2\nexit\n")
	var out bytes.Buffer
	var errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	got := out.String()
	want := "5\n"
	if got != want {
		t.Errorf("Run(\"10 / 2\\nexit\\n\") out = %q, want %q", got, want)
	}
}

// TestRun_ExitStopsProcessing verifies that after 'exit', no further expressions
// in the input are processed and nothing extra is written to out.
func TestRun_ExitStopsProcessing(t *testing.T) {
	in := strings.NewReader("exit\n3 + 4\n")
	var out bytes.Buffer
	var errOut bytes.Buffer

	repl.Run(in, &out, &errOut)

	if out.Len() != 0 {
		t.Errorf("Run: after exit, out = %q, want empty (no further processing)", out.String())
	}
}
