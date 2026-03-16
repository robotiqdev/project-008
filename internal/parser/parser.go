package parser

import (
	"strings"

	calcerrors "github.com/repo/calculator/internal/errors"
)

// Command holds the parsed tokens from a single calculator input line.
type Command struct {
	Op string
	A  string
	B  string
}

// Tokenize splits a raw input line into tokens.
// It returns a single-element slice for "exit" or "quit", a 3-element slice for
// valid "A op B" expressions, or nil and ErrInvalidTokenCount otherwise.
func Tokenize(line string) ([]string, error) {
	tokens := strings.Fields(line)
	if len(tokens) == 1 && (tokens[0] == "exit" || tokens[0] == "quit") {
		return tokens, nil
	}
	if len(tokens) != 3 {
		return nil, calcerrors.ErrInvalidTokenCount
	}
	return tokens, nil
}

// IsExitCommand reports whether cmd is an exit or quit command.
func IsExitCommand(cmd Command) bool {
	panic("not implemented")
}

// ParseCommand parses a slice of tokens into a Command.
// For a single token (e.g. "exit"), only Op is set.
// For three tokens (infix: A op B), all fields are set with Op normalized to lowercase.
func ParseCommand(tokens []string) (Command, error) {
	if len(tokens) == 1 {
		return Command{Op: tokens[0]}, nil
	}
	return Command{Op: strings.ToLower(tokens[1]), A: tokens[0], B: tokens[2]}, nil
}
