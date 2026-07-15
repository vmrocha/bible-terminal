package cli

import (
	"context"
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vmrocha/bible-terminal/internal/bible"
	"github.com/vmrocha/bible-terminal/internal/reference"
	"github.com/vmrocha/bible-terminal/internal/render"
)

// PassageReader is the read behavior required by the CLI.
type PassageReader interface {
	Read(context.Context, reference.Query) (bible.Passage, error)
	Close() error
}

// ReaderFactory opens the offline Bible reader on demand.
type ReaderFactory func(context.Context) (PassageReader, error)

type configuration struct {
	readerFactory ReaderFactory
}

// Option configures optional root-command dependencies.
type Option func(*configuration)

// WithReaderFactory enables commands that access the embedded Bible.
func WithReaderFactory(factory ReaderFactory) Option {
	return func(configuration *configuration) {
		configuration.readerFactory = factory
	}
}

func newReadCommand(factory ReaderFactory) *cobra.Command {
	var plain bool
	command := &cobra.Command{
		Use:     "read <reference>",
		Short:   "Read a chapter, verse, or verse range",
		Example: "  bible read John 3\n  bible read 'John 3:16'\n  bible read 'John 3:16-21' --plain",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(command *cobra.Command, args []string) error {
			query, err := reference.Parse(strings.Join(args, " "))
			if err != nil {
				return err
			}
			if factory == nil {
				return errors.New("Bible reader is unavailable")
			}

			reader, err := factory(command.Context())
			if err != nil {
				return err
			}
			passage, readErr := reader.Read(command.Context(), query)
			closeErr := reader.Close()
			if readErr != nil {
				return readErr
			}
			if closeErr != nil {
				return closeErr
			}

			return render.Passage(command.OutOrStdout(), passage, plain)
		},
	}
	command.Flags().BoolVar(&plain, "plain", false, "emit stable tab-separated verse lines")
	return command
}
