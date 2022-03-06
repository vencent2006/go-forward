package main

import (
	"context"
	"go-examples/course/go-im-cloud/chatdemo/client"
	"go-examples/course/go-im-cloud/chatdemo/server"

	"github.com/sirupsen/logrus"

	"flag"

	"github.com/spf13/cobra"
)

const version = "v1"

func main() {
	flag.Parse()
	// rootCmd
	root := &cobra.Command{
		Use:     "chat",
		Version: version,
		Short:   "chat demo",
	}

	ctx := context.Background()
	root.AddCommand(client.RunCmd(ctx))
	root.AddCommand(server.RunCmd(ctx))

	err := root.Execute()
	if err != nil {
		logrus.WithError(err).Fatal("could not run command")
	}
}
