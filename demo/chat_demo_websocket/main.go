package main

import (
	"context"
	"go-examples/demo/chat_demo_websocket/client"
	"go-examples/demo/chat_demo_websocket/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//func main() {
//	flag.Parse()
//	root := &cobra.Command{
//		Use:     "niubi",
//		Version: "v1",
//		Short:   "niubi demo",
//	}
//	ctx := context.Background()
//	root.AddCommand(server.NewServerStartCmd(ctx))
//
//	if err := root.Execute(); err != nil {
//		logrus.WithError(err).Fatal("cannot run command")
//	}
//}

const version = "v1"

var (
	pid int
)

func main() {
	//flag.Parse()

	root := &cobra.Command{
		Use:     "chat",
		Version: version,
		Short:   "chat demo",
	}
	ctx := context.Background()

	root.AddCommand(server.NewServerStartCmd(ctx))
	root.AddCommand(client.NewClientCmd(ctx))

	if err := root.Execute(); err != nil {
		logrus.WithError(err).Fatal("Could not run command")
	}
}
