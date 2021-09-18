/**
 * @Author: vincent
 * @Description:
 * @File:  root
 * @Version: 1.0.0
 * @Date: 2021/7/25 19:10
 */

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	//rootCmd.AddCommand(timeCmd)
	//rootCmd.AddCommand(jsonCmd)
	//rootCmd.AddCommand(sqlCmd)
}
