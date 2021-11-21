/**
 * @Author: vincent
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2021/10/29 07:22
 */

package console

import (
	"go-examples/course/handwriting-web-inf/code_28/app/console/command/demo"
	"go-examples/course/handwriting-web-inf/code_28/framework"
	"go-examples/course/handwriting-web-inf/code_28/framework/cobra"
	"go-examples/course/handwriting-web-inf/code_28/framework/command"
)

// RunCommand 初始化根Command并运行
func RunCommand(container framework.Container) error {
	// 根Command
	var rootCmd = &cobra.Command{
		// 定义根命令的关键字
		Use: "hade",
		// 简短介绍
		Short: "hade 命令",
		// 根命令的详细介绍
		Long: "hade 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令, 也能很方便编写业务命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	// 为跟Command设置服务容器
	rootCmd.SetContainer(container)
	// 绑定框架的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)

	// 执行rootCommand
	return rootCmd.Execute()
}

// 绑定业务的命令
func AddAppCommand(rootCmd *cobra.Command) {
	// demo例子
	rootCmd.AddCommand(demo.InitFoo())

	// 每秒调用一次Foo命令
	rootCmd.AddCronCommand("* * * * * *", demo.FooCommand)

	// 启动一个分布式任务调度，调度的服务名称为init_func_for_test，每个节点每5s调用一次Foo命令，抢占到了调度任务的节点将抢占锁持续挂载2s才释放
	//rootCmd.AddDistributedCronCommand("foo_func_for_test", "*/5 * * * * *", demo.FooCommand, 2*time.Second)
}
