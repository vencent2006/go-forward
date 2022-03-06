/**
 * @Author: vincent
 * @Description:
 * @File:  cmd
 * @Version: 1.0.0
 * @Date: 2022/3/1 23:13
 */

package server

import (
	"context"

	"github.com/spf13/cobra"
)

const version = "v1"

type options struct {
	// server id
	id string
	// listen address
	listen string
}

func RunCmd(ctx context.Context) *cobra.Command {
	opts := &options{}
	var cmd = &cobra.Command{
		Use:   "server",
		Short: "chat server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(ctx, opts, version)
		},
	}

	cmd.PersistentFlags().StringVarP(&opts.listen, "listen", "l", ":8000", "listen address")
	cmd.PersistentFlags().StringVarP(&opts.id, "id", "i", "", "server id")

	return cmd
}
