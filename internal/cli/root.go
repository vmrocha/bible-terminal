package cli

import (
	"github.com/spf13/cobra"
	"github.com/vmrocha/bible-terminal/internal/buildinfo"
)

// New constructs the root command with the supplied build metadata.
func New(info buildinfo.Info, options ...Option) *cobra.Command {
	configuration := configuration{}
	for _, option := range options {
		option(&configuration)
	}

	command := &cobra.Command{
		Use:           "bible",
		Short:         "Read the Bible from your terminal",
		SilenceErrors: true,
		SilenceUsage:  true,
		Version:       info.Version,
	}

	command.SetVersionTemplate("bible {{.Version}}\n")
	command.AddCommand(newBooksCommand())
	command.AddCommand(newReadCommand(configuration.readerFactory))
	command.AddCommand(newVersionCommand(info))

	return command
}
