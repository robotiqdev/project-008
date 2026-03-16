package repl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	calcerrors "github.com/repo/calculator/internal/errors"
	"github.com/repo/calculator/internal/calc"
	"github.com/repo/calculator/internal/numparse"
	"github.com/repo/calculator/internal/parser"
)

// Run reads lines from r, writing output to out and errors to errOut.
func Run(r io.Reader, out io.Writer, errOut io.Writer) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		tokens, err := parser.Tokenize(line)
		if err != nil {
			calcerrors.HandleError(err, errOut)
			continue
		}

		cmd, err := parser.ParseCommand(tokens)
		if err != nil {
			calcerrors.HandleError(err, errOut)
			continue
		}

		if parser.IsExitCommand(cmd) {
			return
		}

		a, err := numparse.ParseNumber(cmd.A)
		if err != nil {
			calcerrors.HandleError(err, errOut)
			continue
		}

		b, err := numparse.ParseNumber(cmd.B)
		if err != nil {
			calcerrors.HandleError(err, errOut)
			continue
		}

		result, err := calc.Calculate(cmd.Op, a, b)
		if err != nil {
			calcerrors.HandleError(err, errOut)
			continue
		}

		fmt.Fprintln(out, formatResult(result))
	}
}

func formatResult(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
