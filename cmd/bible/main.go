package main

import (
	"fmt"
	"os"

	"github.com/vmrocha/bible-terminal/internal/buildinfo"
	"github.com/vmrocha/bible-terminal/internal/cli"
)

func main() {
	command := cli.New(buildinfo.Current())
	if err := command.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
