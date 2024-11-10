package cli

import (
	"os"

	command "github.com/sanatan01/hither/cli/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hither",
	Short: "Hither is a simple wrapper around containerd",
	Long:  `usage: hither [OPTIONS] COMMAND`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	subCommands := command.GetCommands()
	for _, cmd := range subCommands {
		rootCmd.AddCommand(cmd)
	}
}
