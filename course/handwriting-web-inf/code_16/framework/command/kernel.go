/**
 * @Author: vincent
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2021/10/29 10:07
 */

package command

import "go-examples/course/handwriting-web-inf/code_16/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(DemoCommand)       // demo 命令
	root.AddCommand(initAppCommand())  // app 命令
	root.AddCommand(initEnvCommand())  // env 命令
	root.AddCommand(initCronCommand()) // cron 命令
}
