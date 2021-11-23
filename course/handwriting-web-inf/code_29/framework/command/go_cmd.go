/**
 * @Author: vincent
 * @Description:
 * @File:  go_cmd
 * @Version: 1.0.0
 * @Date: 2021/11/7 12:19
 */

package command

import (
	"go-examples/course/handwriting-web-inf/code_29/framework/cobra"
	"log"
	"os"
	"os/exec"
)

// go just run local go bin
var goCommand = &cobra.Command{
	Use:   "go",
	Short: "运行path/go程序，要求go 必须安装",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("go")
		if err != nil {
			log.Fatalln("hade go: should install go in your PATH")
		}

		cmd := exec.Command(path, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return nil
	},
}
