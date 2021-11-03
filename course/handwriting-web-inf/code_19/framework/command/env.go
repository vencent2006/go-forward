/**
 * @Author: vincent
 * @Description:
 * @File:  env
 * @Version: 1.0.0
 * @Date: 2021/11/1 11:11
 */

package command

import (
	"fmt"
	"go-examples/course/handwriting-web-inf/code_19/framework/cobra"
	"go-examples/course/handwriting-web-inf/code_19/framework/contract"
	"go-examples/course/handwriting-web-inf/code_19/framework/util"
)

// initEnvCommand 获取env相关的命令
func initEnvCommand() *cobra.Command {
	envCommand.AddCommand(envListCommand)
	return envCommand
}

// envCommand 获取当前的App环境
var envCommand = &cobra.Command{
	Use:   "env",
	Short: "获取当前的App环境",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取env环境
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		// 打印环境
		fmt.Println("environment: ", envService.AppEnv())
	},
}

// envListCommand 获取所有的App环境变量
var envListCommand = &cobra.Command{
	Use:   "list",
	Short: "获取所有的环境变量",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取env环境
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		envs := envService.All()
		outs := [][]string{}
		for k, v := range envs {
			outs = append(outs, []string{k, v})
		}
		util.PrettyPrint(outs)
	},
}
