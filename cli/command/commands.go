package command

import (
	"github.com/spf13/cobra"
)

// AddAllCommands Adds all the commands from cli/command to the root command
func AddAllCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		NewDropCommand(),
		NewGetCommand(),
	)
}
