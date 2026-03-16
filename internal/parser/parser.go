package parser

import (
	"fmt"
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
	_ = calcerrors.ErrInvalidTokenCount // stub: not yet implemented
	return nil, nil
}

// ParseCommand parses a raw input line into a Command.
// Expected format: "A op B" (space-separated tokens).
func ParseCommand(line string) (Command, error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return Command{}, fmt.Errorf("invalid input: expected 3 tokens, got %d", len(fields))
	}
	return Command{A: fields[0], Op: fields[1], B: fields[2]}, nil
}
