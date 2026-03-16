package repl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/repo/calculator/internal/calc"
)

// Run reads expressions from in, evaluates them, and writes results to out.
// Errors are written to errOut. Run returns when "exit" is read from in.
func Run(in io.Reader, out io.Writer, errOut io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "exit" {
			return
		}
		parts := strings.Fields(line)
		if len(parts) != 3 {
			fmt.Fprintf(errOut, "Error: invalid expression %q\n", line)
			continue
		}
		a, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Fprintf(errOut, "Error: invalid operand %q\n", parts[0])
			continue
		}
		b, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Fprintf(errOut, "Error: invalid operand %q\n", parts[2])
			continue
		}
		result, err := calc.Calculate(parts[1], a, b)
		if err != nil {
			fmt.Fprintf(errOut, "Error: %v\n", err)
			continue
		}
		fmt.Fprintln(out, formatResult(result))
	}
}

func formatResult(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
