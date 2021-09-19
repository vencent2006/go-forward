/**
 * @Author: vincent
 * @Description:
 * @File:  cmd
 * @Version: 1.0.0
 * @Date: 2021/8/14 22:10
 */

package mock

import (
	"context"

	"github.com/segmentio/ksuid"

	"github.com/spf13/cobra"
)

type StartOptions struct {
	addr     string
	protocol string
}

// client cmd
func NewClientCmd(ctx context.Context) *cobra.Command {
	opts := &StartOptions{}

	cmd := &cobra.Command{
		Use:   "mock_cli",
		Short: "start client",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runClient(ctx, opts)
		},
	}

	cmd.PersistentFlags().StringVarP(&opts.addr, "address", "a", "ws://localhost:8000", "server address")
	cmd.PersistentFlags().StringVarP(&opts.protocol, "protocol", "p", "ws", "protocol ws or tcp")

	return cmd
}

func runClient(ctx context.Context, opts *StartOptions) error {
	client := ClientDemo{}
	client.Start(ksuid.New().String(), opts.protocol, opts.addr)
	return nil
}

func NewServerCmd(ctx context.Context) *cobra.Command {
	opts := &StartOptions{}

	cmd := &cobra.Command{
		Use:   "mock_srv",
		Short: "start server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(ctx, opts)
		},
	}

	cmd.PersistentFlags().StringVarP(&opts.addr, "address", "a", ":8000", "listen address")
	cmd.PersistentFlags().StringVarP(&opts.protocol, "protocol", "p", "ws", "protocol ws or tcp")

	return cmd
}

func runServer(ctx context.Context, opts *StartOptions) error {
	srv := ServerDemo{}
	srv.Start("svr1", opts.protocol, opts.addr)
	return nil
}
