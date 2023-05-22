package main

import (
	"dmitysh/dropper/cli/command"
	"dmitysh/dropper/configs/config"
	"github.com/spf13/cobra"
	"os"
)

const cfgPath = "configs/app.env"

func main() {
	if runDropperErr := runDropper(); runDropperErr != nil {
		os.Exit(1)
	}
}

func runDropper() error {
	config.LoadEnvConfig(cfgPath)
	topCmd := newDropperCommand()
	return topCmd.Execute()
}

func newDropperCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dropper",
		Short: "Tool for fast file drop between machines in local net",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	command.AddAllCommands(cmd)

	return cmd
}
