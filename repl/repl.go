package repl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/repo/calculator/internal/calc"
	calcerrors "github.com/repo/calculator/internal/errors"
)

// Run reads lines from in, evaluates each as "<number> <operator> <number>",
// writes results to out, and writes error messages to errOut.
// It continues processing subsequent lines after errors.
func Run(in io.Reader, out io.Writer, errOut io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		tokens := strings.Fields(line)
		if len(tokens) != 3 {
			calcerrors.HandleError(calcerrors.ErrInvalidTokenCount, errOut)
			continue
		}
		a, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			calcerrors.HandleError(calcerrors.ErrInvalidInput, errOut)
			continue
		}
		b, err := strconv.ParseFloat(tokens[2], 64)
		if err != nil {
			calcerrors.HandleError(calcerrors.ErrInvalidInput, errOut)
			continue
		}
		result, err := calc.Calculate(tokens[1], a, b)
		if err != nil {
			calcerrors.HandleError(err, errOut)
			continue
		}
		fmt.Fprintln(out, result)
	}
}
