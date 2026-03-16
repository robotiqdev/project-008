package parser_test

import (
	"reflect"
	"testing"

	calcerrors "github.com/repo/calculator/internal/errors"
	"github.com/repo/calculator/internal/parser"
)

// TestCommandStruct_HasOpField verifies that Command has an Op field of type string.
func TestCommandStruct_HasOpField(t *testing.T) {
	var cmd parser.Command
	rt := reflect.TypeOf(cmd)

	field, ok := rt.FieldByName("Op")
	if !ok {
		t.Fatal("Command struct is missing field Op")
	}
	if field.Type.Kind() != reflect.String {
		t.Errorf("Command.Op type = %v, want string", field.Type.Kind())
	}
}

// TestCommandStruct_HasAField verifies that Command has an A field of type string.
func TestCommandStruct_HasAField(t *testing.T) {
	var cmd parser.Command
	rt := reflect.TypeOf(cmd)

	field, ok := rt.FieldByName("A")
	if !ok {
		t.Fatal("Command struct is missing field A")
	}
	if field.Type.Kind() != reflect.String {
		t.Errorf("Command.A type = %v, want string", field.Type.Kind())
	}
}

// TestCommandStruct_HasBField verifies that Command has a B field of type string.
func TestCommandStruct_HasBField(t *testing.T) {
	var cmd parser.Command
	rt := reflect.TypeOf(cmd)

	field, ok := rt.FieldByName("B")
	if !ok {
		t.Fatal("Command struct is missing field B")
	}
	if field.Type.Kind() != reflect.String {
		t.Errorf("Command.B type = %v, want string", field.Type.Kind())
	}
}

// TestCommandStruct_ExactlyThreeFields verifies that Command has exactly three exported fields:
// Op, A, and B — no more, no less.
func TestCommandStruct_ExactlyThreeFields(t *testing.T) {
	var cmd parser.Command
	rt := reflect.TypeOf(cmd)

	var exportedFields []string
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if f.IsExported() {
			exportedFields = append(exportedFields, f.Name)
		}
	}

	expected := []string{"Op", "A", "B"}
	if len(exportedFields) != len(expected) {
		t.Errorf("Command has %d exported fields %v, want exactly %d: %v",
			len(exportedFields), exportedFields, len(expected), expected)
		return
	}
	for i, name := range expected {
		if exportedFields[i] != name {
			t.Errorf("Command exported field[%d] = %q, want %q", i, exportedFields[i], name)
		}
	}
}

// TestCommandStruct_FieldsAreStrings verifies all three fields (Op, A, B) are of kind string.
func TestCommandStruct_FieldsAreStrings(t *testing.T) {
	var cmd parser.Command
	rt := reflect.TypeOf(cmd)

	for _, name := range []string{"Op", "A", "B"} {
		field, ok := rt.FieldByName(name)
		if !ok {
			t.Errorf("Command struct is missing field %s", name)
			continue
		}
		if field.Type.Kind() != reflect.String {
			t.Errorf("Command.%s type = %v, want string", name, field.Type.Kind())
		}
	}
}

// TestCommandStruct_ZeroValue verifies the zero value of Command has empty string fields.
func TestCommandStruct_ZeroValue(t *testing.T) {
	cmd := parser.Command{}
	if cmd.Op != "" {
		t.Errorf("zero-value Command.Op = %q, want empty string", cmd.Op)
	}
	if cmd.A != "" {
		t.Errorf("zero-value Command.A = %q, want empty string", cmd.A)
	}
	if cmd.B != "" {
		t.Errorf("zero-value Command.B = %q, want empty string", cmd.B)
	}
}

// TestCommandStruct_LiteralAssignment verifies fields can be set and read back.
func TestCommandStruct_LiteralAssignment(t *testing.T) {
	cmd := parser.Command{Op: "+", A: "3", B: "5"}
	if cmd.Op != "+" {
		t.Errorf("Command.Op = %q, want %q", cmd.Op, "+")
	}
	if cmd.A != "3" {
		t.Errorf("Command.A = %q, want %q", cmd.A, "3")
	}
	if cmd.B != "5" {
		t.Errorf("Command.B = %q, want %q", cmd.B, "5")
	}
}

// TestTokenize_BasicExpression verifies Tokenize splits a simple "A op B" expression correctly.
func TestTokenize_BasicExpression(t *testing.T) {
	tokens, err := parser.Tokenize("3 + 4")
	if err != nil {
		t.Fatalf("Tokenize(%q) unexpected error: %v", "3 + 4", err)
	}
	expected := []string{"3", "+", "4"}
	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Tokenize(%q) = %v, want %v", "3 + 4", tokens, expected)
	}
}

// TestTokenize_SingleToken verifies Tokenize returns ErrInvalidTokenCount for a single non-command token.
func TestTokenize_SingleToken(t *testing.T) {
	tokens, err := parser.Tokenize("3")
	if tokens != nil {
		t.Errorf("Tokenize(%q) tokens = %v, want nil", "3", tokens)
	}
	if err != calcerrors.ErrInvalidTokenCount {
		t.Errorf("Tokenize(%q) error = %v, want ErrInvalidTokenCount", "3", err)
	}
}

// TestTokenize_TooManyTokens verifies Tokenize returns ErrInvalidTokenCount when more than 3 tokens are provided.
func TestTokenize_TooManyTokens(t *testing.T) {
	tokens, err := parser.Tokenize("3 + 4 5")
	if tokens != nil {
		t.Errorf("Tokenize(%q) tokens = %v, want nil", "3 + 4 5", tokens)
	}
	if err != calcerrors.ErrInvalidTokenCount {
		t.Errorf("Tokenize(%q) error = %v, want ErrInvalidTokenCount", "3 + 4 5", err)
	}
}

// TestTokenize_MultipleSpaces verifies Tokenize handles multiple/mixed whitespace between tokens.
func TestTokenize_MultipleSpaces(t *testing.T) {
	tokens, err := parser.Tokenize("  10  *  2  ")
	if err != nil {
		t.Fatalf("Tokenize(%q) unexpected error: %v", "  10  *  2  ", err)
	}
	expected := []string{"10", "*", "2"}
	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Tokenize(%q) = %v, want %v", "  10  *  2  ", tokens, expected)
	}
}

// TestTokenize_EmptyString verifies Tokenize returns ErrInvalidTokenCount for an empty input.
func TestTokenize_EmptyString(t *testing.T) {
	tokens, err := parser.Tokenize("")
	if tokens != nil {
		t.Errorf("Tokenize(%q) tokens = %v, want nil", "", tokens)
	}
	if err != calcerrors.ErrInvalidTokenCount {
		t.Errorf("Tokenize(%q) error = %v, want ErrInvalidTokenCount", "", err)
	}
}

// TestTokenize_ExitCommand verifies Tokenize returns ["exit"] with no error for the exit command.
func TestTokenize_ExitCommand(t *testing.T) {
	tokens, err := parser.Tokenize("exit")
	if err != nil {
		t.Fatalf("Tokenize(%q) unexpected error: %v", "exit", err)
	}
	expected := []string{"exit"}
	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Tokenize(%q) = %v, want %v", "exit", tokens, expected)
	}
}

// TestTokenize_QuitCommand verifies Tokenize returns ["quit"] with no error for the quit command.
func TestTokenize_QuitCommand(t *testing.T) {
	tokens, err := parser.Tokenize("quit")
	if err != nil {
		t.Fatalf("Tokenize(%q) unexpected error: %v", "quit", err)
	}
	expected := []string{"quit"}
	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Tokenize(%q) = %v, want %v", "quit", tokens, expected)
	}
}

// TestParseCommand_BasicArithmetic verifies ParseCommand(["3", "+", "4"]) returns
// Command{Op:"+", A:"3", B:"4"} with no error.
func TestParseCommand_BasicArithmetic(t *testing.T) {
	tokens := []string{"3", "+", "4"}
	cmd, err := parser.ParseCommand(tokens)
	if err != nil {
		t.Fatalf("ParseCommand(%v) unexpected error: %v", tokens, err)
	}
	if cmd.A != "3" {
		t.Errorf("ParseCommand(%v).A = %q, want %q", tokens, cmd.A, "3")
	}
	if cmd.Op != "+" {
		t.Errorf("ParseCommand(%v).Op = %q, want %q", tokens, cmd.Op, "+")
	}
	if cmd.B != "4" {
		t.Errorf("ParseCommand(%v).B = %q, want %q", tokens, cmd.B, "4")
	}
}

// TestParseCommand_WordOperator verifies ParseCommand(["10", "divide", "2"]) returns
// Command{Op:"divide", A:"10", B:"2"} with no error.
func TestParseCommand_WordOperator(t *testing.T) {
	tokens := []string{"10", "divide", "2"}
	cmd, err := parser.ParseCommand(tokens)
	if err != nil {
		t.Fatalf("ParseCommand(%v) unexpected error: %v", tokens, err)
	}
	if cmd.A != "10" {
		t.Errorf("ParseCommand(%v).A = %q, want %q", tokens, cmd.A, "10")
	}
	if cmd.Op != "divide" {
		t.Errorf("ParseCommand(%v).Op = %q, want %q", tokens, cmd.Op, "divide")
	}
	if cmd.B != "2" {
		t.Errorf("ParseCommand(%v).B = %q, want %q", tokens, cmd.B, "2")
	}
}

// TestParseCommand_ExitCommand verifies ParseCommand(["exit"]) returns
// Command{Op:"exit"} with A and B empty, and no error.
func TestParseCommand_ExitCommand(t *testing.T) {
	tokens := []string{"exit"}
	cmd, err := parser.ParseCommand(tokens)
	if err != nil {
		t.Fatalf("ParseCommand(%v) unexpected error: %v", tokens, err)
	}
	if cmd.Op != "exit" {
		t.Errorf("ParseCommand(%v).Op = %q, want %q", tokens, cmd.Op, "exit")
	}
	if cmd.A != "" {
		t.Errorf("ParseCommand(%v).A = %q, want empty string", tokens, cmd.A)
	}
	if cmd.B != "" {
		t.Errorf("ParseCommand(%v).B = %q, want empty string", tokens, cmd.B)
	}
}

// TestParseCommand_OpIsMiddleToken verifies that the operator is taken from position [1]
// (infix notation: A op B) and A from [0], B from [2].
func TestParseCommand_OpIsMiddleToken(t *testing.T) {
	tokens := []string{"100", "multiply", "5"}
	cmd, err := parser.ParseCommand(tokens)
	if err != nil {
		t.Fatalf("ParseCommand(%v) unexpected error: %v", tokens, err)
	}
	if cmd.A != tokens[0] {
		t.Errorf("ParseCommand(%v).A = %q, want tokens[0] = %q", tokens, cmd.A, tokens[0])
	}
	if cmd.Op != tokens[1] {
		t.Errorf("ParseCommand(%v).Op = %q, want tokens[1] = %q", tokens, cmd.Op, tokens[1])
	}
	if cmd.B != tokens[2] {
		t.Errorf("ParseCommand(%v).B = %q, want tokens[2] = %q", tokens, cmd.B, tokens[2])
	}
}

// TestParseCommand_OperatorNormalizedToLower verifies that uppercase operators are
// normalized to lowercase (e.g. "ADD" -> "add", "DIVIDE" -> "divide").
func TestParseCommand_OperatorNormalizedToLower(t *testing.T) {
	tests := []struct {
		tokens  []string
		wantOp  string
	}{
		{[]string{"3", "ADD", "4"}, "add"},
		{[]string{"10", "DIVIDE", "2"}, "divide"},
		{[]string{"5", "MuLtIpLy", "3"}, "multiply"},
		{[]string{"8", "SUBtract", "1"}, "subtract"},
	}

	for _, tt := range tests {
		cmd, err := parser.ParseCommand(tt.tokens)
		if err != nil {
			t.Fatalf("ParseCommand(%v) unexpected error: %v", tt.tokens, err)
		}
		if cmd.Op != tt.wantOp {
			t.Errorf("ParseCommand(%v).Op = %q, want %q (lowercase)", tt.tokens, cmd.Op, tt.wantOp)
		}
	}
}

// TestParseCommand_QuitCommand verifies ParseCommand(["quit"]) returns
// Command{Op:"quit"} with no error (single-token command like exit).
func TestParseCommand_QuitCommand(t *testing.T) {
	tokens := []string{"quit"}
	cmd, err := parser.ParseCommand(tokens)
	if err != nil {
		t.Fatalf("ParseCommand(%v) unexpected error: %v", tokens, err)
	}
	if cmd.Op != "quit" {
		t.Errorf("ParseCommand(%v).Op = %q, want %q", tokens, cmd.Op, "quit")
	}
	if cmd.A != "" {
		t.Errorf("ParseCommand(%v).A = %q, want empty string", tokens, cmd.A)
	}
	if cmd.B != "" {
		t.Errorf("ParseCommand(%v).B = %q, want empty string", tokens, cmd.B)
	}
}

// TestParseCommand_NoError_ThreeTokens verifies ParseCommand returns nil error for valid 3-token input.
func TestParseCommand_NoError_ThreeTokens(t *testing.T) {
	tokens := []string{"7", "-", "2"}
	_, err := parser.ParseCommand(tokens)
	if err != nil {
		t.Errorf("ParseCommand(%v) error = %v, want nil", tokens, err)
	}
}

// TestParseCommand_NoError_OneToken verifies ParseCommand returns nil error for a single-token command.
func TestParseCommand_NoError_OneToken(t *testing.T) {
	tokens := []string{"exit"}
	_, err := parser.ParseCommand(tokens)
	if err != nil {
		t.Errorf("ParseCommand(%v) error = %v, want nil", tokens, err)
	}
}

// TestIsExitCommand_ExitReturnsTrue verifies IsExitCommand returns true for Command{Op:"exit"}.
func TestIsExitCommand_ExitReturnsTrue(t *testing.T) {
	cmd := parser.Command{Op: "exit"}
	if !parser.IsExitCommand(cmd) {
		t.Errorf("IsExitCommand(%v) = false, want true", cmd)
	}
}

// TestIsExitCommand_QuitReturnsTrue verifies IsExitCommand returns true for Command{Op:"quit"}.
func TestIsExitCommand_QuitReturnsTrue(t *testing.T) {
	cmd := parser.Command{Op: "quit"}
	if !parser.IsExitCommand(cmd) {
		t.Errorf("IsExitCommand(%v) = false, want true", cmd)
	}
}

// TestIsExitCommand_AddReturnsFalse verifies IsExitCommand returns false for Command{Op:"add"}.
func TestIsExitCommand_AddReturnsFalse(t *testing.T) {
	cmd := parser.Command{Op: "add"}
	if parser.IsExitCommand(cmd) {
		t.Errorf("IsExitCommand(%v) = true, want false", cmd)
	}
}

// TestIsExitCommand_PlusReturnsFalse verifies IsExitCommand returns false for Command{Op:"+"}.
func TestIsExitCommand_PlusReturnsFalse(t *testing.T) {
	cmd := parser.Command{Op: "+"}
	if parser.IsExitCommand(cmd) {
		t.Errorf("IsExitCommand(%v) = true, want false", cmd)
	}
}
