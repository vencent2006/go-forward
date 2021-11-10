/**
 * @Author: vincent
 * @Description:
 * @File:  npm
 * @Version: 1.0.0
 * @Date: 2021/11/5 16:58
 */

package command

import (
	"go-examples/course/handwriting-web-inf/code_24/framework/cobra"
	"log"
	"os"
	"os/exec"
)

// npm just run local go bin
var npmCommand = &cobra.Command{
	Use:   "npm",
	Short: "运行 PATH/npm 命令",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("npm")
		if err != nil {
			log.Fatalln("hade npm: should install npm in your PATH")
		}

		cmd := exec.Command(path, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return nil
	},
}
