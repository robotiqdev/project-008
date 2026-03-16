package parser

// Command holds the parsed tokens from a single calculator input line.
type Command struct {
	Op string
	A  string
	B  string
}

// ParseCommand parses a raw input line into a Command.
// This is a stub; the full implementation will be provided in TASK-4271.
func ParseCommand(line string) (Command, error) {
	return Command{}, nil
}
