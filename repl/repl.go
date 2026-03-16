package repl

import "io"

// Run reads lines from in, evaluates each as "<number> <operator> <number>",
// writes results to out, and writes error messages to errOut.
// It continues processing subsequent lines after errors.
func Run(in io.Reader, out io.Writer, errOut io.Writer) {
}
