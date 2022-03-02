/**
 * @Author: vincent
 * @Description:
 * @File:  cmd
 * @Version: 1.0.0
 * @Date: 2022/3/1 23:13
 */

package client

import (
	"context"

	"github.com/spf13/cobra"
)

type Options struct {
	address string
	user    string
}

func RunCmd(ctx context.Context) *cobra.Command {
	// options
	opts := &Options{}

	// cmd
	cmd := &cobra.Command{
		Use:   "client",
		Short: "Start client",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runClient(ctx, opts)
		},
	}

	// flags
	cmd.PersistentFlags().StringVarP(&opts.address, "address", "a", "ws://127.0.0.1:8000", "server address")
	cmd.PersistentFlags().StringVarP(&opts.user, "user", "u", "", "user")

	return cmd
}
