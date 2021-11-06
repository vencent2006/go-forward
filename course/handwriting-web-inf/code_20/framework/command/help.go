package command

import (
	"fmt"

	"go-examples/course/handwriting-web-inf/code_20/framework/cobra"
	"go-examples/course/handwriting-web-inf/code_20/framework/contract"
)

// helpCommand show current envionment
var DemoCommand = &cobra.Command{
	Use:   "demo",
	Short: "demo for framework",
	Run: func(c *cobra.Command, args []string) {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		fmt.Println("app base folder:", appService.BaseFolder())
	},
}
