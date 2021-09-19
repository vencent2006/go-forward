/**
* @Author: vincent
* @Description:
* @File:  start
* @Version: 1.0.0
* @Date: 2021/8/10 16:04
 */

package server

import (
	"context"

	"github.com/spf13/cobra"
)

type ServerStartOptions struct {
	id     string
	listen string
}

func NewServerStartCmd(ctx context.Context) *cobra.Command {

	opts := &ServerStartOptions{}

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Starts a chat server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServerStart(ctx, opts)
		},
	}

	// flag
	cmd.PersistentFlags().StringVarP(&opts.id, "serverid", "i", "demo", "server id")
	cmd.PersistentFlags().StringVarP(&opts.listen, "listen", "l", ":8000", "listen address")

	return cmd
}

func RunServerStart(ctx context.Context, opts *ServerStartOptions) error {
	server := NewServer(opts.id, opts.listen)
	defer server.Shutdown()
	return server.Start()
}
