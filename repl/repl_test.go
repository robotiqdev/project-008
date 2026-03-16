package repl_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/repo/calculator/repl"
)

// TestRun_ContinuesAfterParseError verifies that Run keeps processing lines after a parse
// error. Input "3\nabc\n3 + 4\n" has invalid lines on lines 1 and 2 (single non-exit
// tokens produce ErrInvalidTokenCount), and the REPL must still produce output for
// the valid "3 + 4" expression on line 3.
func TestRun_ContinuesAfterParseError(t *testing.T) {
	input := "3\nabc\n3 + 4\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	if out.Len() == 0 {
		t.Error("Run() wrote nothing to out; expected a result for '3 + 4' after error lines")
	}
	if errOut.Len() == 0 {
		t.Error("Run() wrote nothing to errOut; expected errors for invalid lines '3' and 'abc'")
	}
}

// TestRun_ErrorsGoToErrOut verifies that parse errors are written to errOut, not out.
// Input "3\n" is a single non-exit token and must produce an error in errOut only.
func TestRun_ErrorsGoToErrOut(t *testing.T) {
	input := "3\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	if errOut.Len() == 0 {
		t.Error("Run() wrote nothing to errOut; expected an error for invalid input '3'")
	}
	if out.Len() != 0 {
		t.Errorf("Run() wrote %q to out for invalid input; errors must not go to out", out.String())
	}
}

// TestRun_ValidResultsGoToOut verifies that the result of a valid expression is written
// to out, not errOut. Input "3 + 4\n" should produce a result in out only.
func TestRun_ValidResultsGoToOut(t *testing.T) {
	input := "3 + 4\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	if out.Len() == 0 {
		t.Error("Run() wrote nothing to out; expected a result for valid expression '3 + 4'")
	}
	if errOut.Len() != 0 {
		t.Errorf("Run() wrote %q to errOut for valid input; results must not go to errOut", errOut.String())
	}
}

// TestRun_ValidResult_AdditionValue verifies that "3 + 4" produces the result "7" in out.
func TestRun_ValidResult_AdditionValue(t *testing.T) {
	input := "3 + 4\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	got := strings.TrimSpace(out.String())
	if got != "7" {
		t.Errorf("Run() out = %q, want %q for input '3 + 4'", got, "7")
	}
}

// TestRun_ParseError_InvalidTokenCount_SingleToken verifies that a single non-exit token
// produces exactly one error line in errOut and nothing in out.
func TestRun_ParseError_InvalidTokenCount_SingleToken(t *testing.T) {
	input := "abc\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	if errOut.Len() == 0 {
		t.Error("Run() wrote nothing to errOut; expected an error for invalid token 'abc'")
	}
	if out.Len() != 0 {
		t.Errorf("Run() wrote %q to out for invalid input; errors must not go to out", out.String())
	}
}

// TestRun_ExitCommandStopsLoop verifies that the exit command causes Run to return
// without processing further input. Lines after "exit" must not produce output.
func TestRun_ExitCommandStopsLoop(t *testing.T) {
	input := "exit\n3 + 4\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	if out.Len() != 0 {
		t.Errorf("Run() wrote %q to out after exit; should have stopped before '3 + 4'", out.String())
	}
	if errOut.Len() != 0 {
		t.Errorf("Run() wrote %q to errOut after exit; should have stopped cleanly", errOut.String())
	}
}

// TestRun_QuitCommandStopsLoop verifies that the quit command causes Run to return
// without processing further input, same as exit.
func TestRun_QuitCommandStopsLoop(t *testing.T) {
	input := "quit\n3 + 4\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	if out.Len() != 0 {
		t.Errorf("Run() wrote %q to out after quit; should have stopped before '3 + 4'", out.String())
	}
	if errOut.Len() != 0 {
		t.Errorf("Run() wrote %q to errOut after quit; should have stopped cleanly", errOut.String())
	}
}

// TestRun_EmptyInput_ProducesNoOutput verifies that an empty reader causes Run to return
// immediately without writing to either out or errOut.
func TestRun_EmptyInput_ProducesNoOutput(t *testing.T) {
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(""), &out, &errOut)

	if out.Len() != 0 {
		t.Errorf("Run() wrote %q to out for empty input; expected nothing", out.String())
	}
	if errOut.Len() != 0 {
		t.Errorf("Run() wrote %q to errOut for empty input; expected nothing", errOut.String())
	}
}

// TestRun_MultipleValidExpressions verifies that Run correctly handles multiple valid
// expressions, writing each result to out on its own line.
func TestRun_MultipleValidExpressions(t *testing.T) {
	input := "2 + 3\n10 - 4\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	if errOut.Len() != 0 {
		t.Errorf("Run() wrote errors %q to errOut for valid inputs", errOut.String())
	}

	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 {
		t.Errorf("Run() produced %d result lines, want 2; out = %q", len(lines), out.String())
		return
	}
	if lines[0] != "5" {
		t.Errorf("Run() first result = %q, want %q for '2 + 3'", lines[0], "5")
	}
	if lines[1] != "6" {
		t.Errorf("Run() second result = %q, want %q for '10 - 4'", lines[1], "6")
	}
}

// TestRun_ErrorOnSecondLine_ResultOnThirdLine is the primary scenario from the task spec.
// Input "3\nabc\n3 + 4\n": lines 1 and 2 are invalid (single non-exit tokens),
// line 3 is valid. The loop must continue past errors and produce "7" in out.
func TestRun_ErrorOnSecondLine_ResultOnThirdLine(t *testing.T) {
	input := "3\nabc\n3 + 4\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	// Both errOut and out must have content.
	if errOut.Len() == 0 {
		t.Error("Run() wrote nothing to errOut; expected errors for '3' and 'abc'")
	}
	if out.Len() == 0 {
		t.Error("Run() wrote nothing to out; expected a result for '3 + 4'")
	}

	// The result of "3 + 4" must be "7".
	got := strings.TrimSpace(out.String())
	if got != "7" {
		t.Errorf("Run() out = %q, want %q for '3 + 4' after error lines", got, "7")
	}
}

// TestRun_ErrorOutputNotInOut verifies that when only errors occur, nothing is written to out.
func TestRun_ErrorOutputNotInOut(t *testing.T) {
	input := "bad input here is too many tokens\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	if out.Len() != 0 {
		t.Errorf("Run() wrote %q to out for invalid input; errors must not appear in out", out.String())
	}
}

// TestRun_ErrorOutputContainsMessage verifies that error output from errOut is non-empty
// and contains meaningful content (not just a newline).
func TestRun_ErrorOutputContainsMessage(t *testing.T) {
	input := "xyz\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	errMsg := errOut.String()
	if len(strings.TrimSpace(errMsg)) == 0 {
		t.Error("Run() errOut contains only whitespace; expected a meaningful error message")
	}
}

// TestRun_BlankLinesSkipped verifies that blank lines produce no output in either out or errOut.
func TestRun_BlankLinesSkipped(t *testing.T) {
	input := "\n\n\n"
	var out, errOut bytes.Buffer

	repl.Run(strings.NewReader(input), &out, &errOut)

	if out.Len() != 0 {
		t.Errorf("Run() wrote %q to out for blank lines; expected nothing", out.String())
	}
	if errOut.Len() != 0 {
		t.Errorf("Run() wrote %q to errOut for blank lines; expected nothing", errOut.String())
	}
}
