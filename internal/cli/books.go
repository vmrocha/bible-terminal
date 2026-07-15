package cli

import (
	"github.com/spf13/cobra"
	"github.com/vmrocha/bible-terminal/internal/canon"
	"github.com/vmrocha/bible-terminal/internal/render"
)

func newBooksCommand() *cobra.Command {
	var plain bool
	command := &cobra.Command{
		Use:   "books",
		Short: "List books and accepted aliases",
		Args:  cobra.NoArgs,
		RunE: func(command *cobra.Command, _ []string) error {
			return render.Books(command.OutOrStdout(), canon.ProtestantBooks(), plain)
		},
	}
	command.Flags().BoolVar(&plain, "plain", false, "emit stable tab-separated book lines")
	return command
}
