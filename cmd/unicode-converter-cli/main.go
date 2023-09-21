package main

import (
	"github.com/martirosharutyunyan/unicode-converter/pkg/cmd"
	"os"
)

func main() {
	cmd.Execute(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
}
