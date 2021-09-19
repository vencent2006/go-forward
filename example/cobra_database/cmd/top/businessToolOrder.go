/**
 * @Author: vincent
 * @Description:
 * @File:  top
 * @Version: 1.0.0
 * @Date: 2021/9/13 14:35
 */

package top

import (
	"github.com/spf13/cobra"
)

var uid string
var pid string
var status string

var businessToolOrderCmd = &cobra.Command{
	Use:   "business_tool_order",
	Short: "use business_tool_order table",
	RunE: func(cmd *cobra.Command, args []string) error {
		return doCommand()
	},
}

func init() {
	businessToolOrderCmd.Flags().StringVarP(&uid, "uid", "u", "", "请输入 uid")
	businessToolOrderCmd.Flags().StringVarP(&pid, "pid", "p", "", "请输入 pid")
	businessToolOrderCmd.Flags().StringVarP(&status, "status", "s", "", "请输入 status")
}

func doCommand() error {
	return nil
}
