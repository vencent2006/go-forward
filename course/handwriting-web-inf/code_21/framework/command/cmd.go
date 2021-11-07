/**
 * @Author: vincent
 * @Description:
 * @File:  cmd
 * @Version: 1.0.0
 * @Date: 2021/11/6 22:37
 */

package command

import (
	"fmt"
	"go-examples/course/handwriting-web-inf/code_21/framework/cobra"
	"go-examples/course/handwriting-web-inf/code_21/framework/contract"
	"go-examples/course/handwriting-web-inf/code_21/framework/util"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/jianfengye/collection"

	"github.com/AlecAivazis/survey/v2"
)

// 初始化command相关命令
func initCmdCommand() *cobra.Command {
	cmdCommand.AddCommand(cmdListCommand)
	cmdCommand.AddCommand(cmdCreateCommand)
	return cmdCommand
}

// 二级命令
var cmdCommand = &cobra.Command{
	Use:   "command",
	Short: "控制台相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

// cmdListCommand 列出所有的控制台命令
var cmdListCommand = &cobra.Command{
	Use:   "list",
	Short: "列出所有控制台命令",
	RunE: func(c *cobra.Command, args []string) error {
		cmds := c.Root().Commands()
		var ps [][]string
		for _, cmd := range cmds {
			line := []string{cmd.Name(), cmd.Short}
			ps = append(ps, line)
		}
		util.PrettyPrint(ps)
		return nil
	},
}

// cmdCreateCommand 创建一个业务控制命令
var cmdCreateCommand = &cobra.Command{
	Use:     "new",
	Aliases: []string{"create", "init"}, // 设置别名为 create init
	Short:   "创建一个控制台命令",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()

		fmt.Println("开始创建控制台命令...")
		var name string
		var folder string

		// todo: 这两个局部代码块的用法，值得学习和记忆
		{
			prompt := &survey.Input{
				Message: "请输入控制台命令名称:",
			}
			err := survey.AskOne(prompt, &name)
			if err != nil {
				return err
			}
		}
		{
			prompt := &survey.Input{
				Message: "请输入文件夹名称(默认:同控制台命令):",
			}
			err := survey.AskOne(prompt, &folder)
			if err != nil {
				return err
			}
		}

		if folder == "" {
			// 目录没有输入，就和name一样
			folder = name
		}

		// 判断文件不存在
		app := container.MustMake(contract.AppKey).(contract.App)

		pFolder := app.CommandFolder()
		fmt.Println("pFolder: ", pFolder)
		fmt.Println("folder: ", folder)
		subFolders, err := util.SubDir(pFolder)
		if err != nil {
			return err
		}
		// todo: 这是作者的一个库，github.com/jianfengye/collection
		subColl := collection.NewStrCollection(subFolders)
		if subColl.Contains(folder) {
			fmt.Println("目录名称已经存在")
			return nil
		}

		// 开始创建文件
		if err := os.Mkdir(filepath.Join(pFolder, folder), 0700); err != nil {
			return err
		}

		// 创建title这个模板方法
		funcs := template.FuncMap{"title": strings.Title}
		{
			// 创建 name.go
			file := filepath.Join(pFolder, folder, name+".go")
			f, err := os.Create(file)
			if err != nil {
				// todo: errors.Cause , from pkg/errors
				return errors.Cause(err)
			}

			// 使用contractTmp模板来初始化template，并且让这个模板支持title方法，即支持{{.|title}}
			t := template.Must(template.New("cmd").Funcs(funcs).Parse(cmdTmpl))
			// 将name传递进入到template中渲染，并且输出到contract.go中
			if err := t.Execute(f, name); err != nil {
				// todo: 使用errors.Cause
				return errors.Cause(err)
			}
		}
		fmt.Println("创建新命令行工具成功，路径:", filepath.Join(pFolder, folder))
		fmt.Println("请记得开发完成后将命令行工具挂载到 console/kernel.go")
		return nil
	},
}

// 命令行工具模版
var cmdTmpl string = `package {{.}}

import (
	"fmt"

	"github.com/gohade/hade/framework/cobra"
)

var {{.|title}}Command = &cobra.Command{
	Use:   "{{.}}",
	Short: "{{.}}",
	RunE: func(c *cobra.Command, args []string) error {
        container := c.GetContainer()
		fmt.Println(container)
		return nil
	},
}

`
