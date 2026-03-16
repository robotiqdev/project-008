package parser

import (
	"fmt"
	"strings"
)

// Command holds the parsed tokens from a single calculator input line.
type Command struct {
	Op string
	A  string
	B  string
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
