package main

import (
	"context"
	"flag"
	"go-examples/chat_demo_im/examples/mock"
	"go-examples/chat_demo_im/logger"

	"github.com/spf13/cobra"
)

const version = "v1"

func main() {
	// todo: 感觉这个flag也没啥用
	flag.Parse()

	// root Cmd
	root := &cobra.Command{
		Use:     "fim",
		Version: version,
		Short:   "server",
	}
	ctx := context.Background()

	// mock: add command
	root.AddCommand(mock.NewClientCmd(ctx))
	root.AddCommand(mock.NewServerCmd(ctx))

	// cmd execute
	if err := root.Execute(); err != nil {
		logger.WithError(err).Fatal("could not run command")
	}
}
