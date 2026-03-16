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

// TestParseCommand_StubExists is a stub test that confirms ParseCommand is callable.
// The full behavioral tests will be added in TASK-4271.
func TestParseCommand_StubExists(t *testing.T) {
	// ParseCommand must exist and accept a string, returning (Command, error).
	// This stub verifies the function signature compiles correctly.
	cmd, err := parser.ParseCommand("3 + 5")

	// With a full implementation, parsing "3 + 5" should produce
	// Command{A: "3", Op: "+", B: "5"} with no error.
	// The current stub returns an empty Command, so this test is expected to fail.
	if err != nil {
		t.Fatalf("ParseCommand(%q) unexpected error: %v", "3 + 5", err)
	}
	if cmd.A != "3" {
		t.Errorf("ParseCommand(%q).A = %q, want %q", "3 + 5", cmd.A, "3")
	}
	if cmd.Op != "+" {
		t.Errorf("ParseCommand(%q).Op = %q, want %q", "3 + 5", cmd.Op, "+")
	}
	if cmd.B != "5" {
		t.Errorf("ParseCommand(%q).B = %q, want %q", "3 + 5", cmd.B, "5")
	}
}
