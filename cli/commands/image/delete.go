package image

import (
	"context"
	"fmt"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/sanatan01/hither/pkg"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete an image",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx, err := pkg.NewContainerdClient()
		if err != nil {
			fmt.Println("Error creating containerd client:", err)
			return
		}
		err = Delete(client, ctx, args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Image deleted successfully")
	},
}

func Delete(client *containerd.Client, ctx context.Context, img string) error {
	_, err := client.GetImage(ctx, img)
	if err != nil {
		return err
	}
	return client.ImageService().Delete(ctx, img)
}
