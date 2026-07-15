package main

import (
	"context"
	"fmt"
	"os"

	"github.com/vmrocha/bible-terminal/internal/buildinfo"
	"github.com/vmrocha/bible-terminal/internal/cli"
	"github.com/vmrocha/bible-terminal/internal/storage"
)

func main() {
	command := cli.New(
		buildinfo.Current(),
		cli.WithReaderFactory(func(ctx context.Context) (cli.PassageReader, error) {
			return storage.OpenEmbedded(ctx)
		}),
	)
	if err := command.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
