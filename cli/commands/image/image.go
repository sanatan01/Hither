package image

import "github.com/spf13/cobra"

func NewImageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Manage images",
	}

	cmd.AddCommand(listCmd)
	cmd.AddCommand(deleteCmd)
	return cmd
}
