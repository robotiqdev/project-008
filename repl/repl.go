package repl

import (
	"bufio"
	"io"
	"strings"
)

// Run reads lines from r, writing output to out and errors to errOut.
func Run(r io.Reader, out io.Writer, errOut io.Writer) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		// TODO: parse and dispatch
	}
}
