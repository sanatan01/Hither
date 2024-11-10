package container

import (
	"fmt"

	"github.com/sanatan01/hither/pkg"
	"github.com/sanatan01/hither/pkg/command/container"
	"github.com/spf13/cobra"
)

// TOOD: Take a container-name opt
var runCmd = &cobra.Command{
	Use:   "run IMAGE",
	Short: "Creates and runs a container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runContainer(args[0])
	},
}

func runContainer(img string) {
	client, ctx, err := pkg.NewContainerdClient()
	if err != nil {
		fmt.Println("Error creating containerd client:", err)
		return
	}
	err = container.Run(client, ctx, img)
	if err != nil {
		fmt.Printf("Error running container: %v\n", err)
	}
}
