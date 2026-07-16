package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCompletionCommand() *cobra.Command {
	command := &cobra.Command{
		Use:       "completion <shell>",
		Short:     "Generate shell completion scripts",
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
		RunE: func(command *cobra.Command, args []string) error {
			writer := command.OutOrStdout()
			root := command.Root()
			switch args[0] {
			case "bash":
				return root.GenBashCompletionV2(writer, true)
			case "zsh":
				return root.GenZshCompletion(writer)
			case "fish":
				return root.GenFishCompletion(writer, true)
			case "powershell":
				return root.GenPowerShellCompletion(writer)
			default:
				return fmt.Errorf("unsupported shell %q", args[0])
			}
		},
	}
	command.DisableFlagsInUseLine = true
	return command
}
