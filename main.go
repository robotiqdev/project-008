package main

import (
	"os"

	"github.com/repo/calculator/repl"
)

func main() {
	repl.Run(os.Stdin, os.Stdout, os.Stderr)
}
