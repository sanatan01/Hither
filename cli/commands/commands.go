package command

import (
	"github.com/sanatan01/hither/cli/commands/container"
	"github.com/sanatan01/hither/cli/commands/image"

	"github.com/spf13/cobra"
)

func GetCommands() []*cobra.Command {
	commands := []*cobra.Command{
		pullCmd,
		image.NewImageCommand(),
		container.NewContainerCommand(),
	}
	return commands
}
