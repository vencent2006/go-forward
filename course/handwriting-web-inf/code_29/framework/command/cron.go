/**
 * @Author: vincent
 * @Description:
 * @File:  cron
 * @Version: 1.0.0
 * @Date: 2021/10/30 21:21
 */

package command

import (
	"fmt"
	"go-examples/course/handwriting-web-inf/code_29/framework/cobra"
	"go-examples/course/handwriting-web-inf/code_29/framework/contract"
	"go-examples/course/handwriting-web-inf/code_29/framework/util"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/sevlyar/go-daemon"
)

var cronDaemon = false

func initCronCommand() *cobra.Command {
	cronStartCommand.Flags().BoolVarP(&cronDaemon, "daemon", "d", false, "start serve daemon")
	cronCommand.AddCommand(cronStartCommand)
	cronCommand.AddCommand(cronRestartCommand)
	cronCommand.AddCommand(cronStopCommand)
	cronCommand.AddCommand(cronListCommand)
	cronCommand.AddCommand(cronStateCommand)
	return cronCommand
}

var cronCommand = &cobra.Command{
	Use:   "cron",
	Short: "定时任务相关命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

// cron进程的启动服务
var cronStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 获取容器
		container := cmd.GetContainer()
		// 获取容器中的app服务
		appService := container.MustMake(contract.AppKey).(contract.App)

		// 设置cron的日志地址和进程id（process id）地址
		pidFolder := appService.RuntimeFolder()
		serverPidFile := filepath.Join(pidFolder, "cron.pid")
		logFolder := appService.LogFolder()
		serverLogFile := filepath.Join(logFolder, "cron.log")
		currentFolder := appService.BaseFolder()

		// daemon模式
		if cronDaemon {
			// 采用daemon模式，创建一个Context
			ctx := &daemon.Context{
				// 设置pid文件
				PidFileName: serverPidFile,
				PidFilePerm: 0664,
				// 设置日志文件
				LogFileName: serverLogFile,
				LogFilePerm: 0640,
				// 设置工作路径
				WorkDir: currentFolder,
				// 设置所有设置文件的mask，默认为750
				Umask: 027, //todo:也不是750呀
				// 子进程的参数，按照这个参数设置，子进程的命令为 ./hade cron start --daemon=true
				Args: []string{"", "cron", "start", "--daemon=true"},
			}

			// 启动子进程，d不为空表示当前是父进程，d为空表示当前是子进程
			d, err := ctx.Reborn()
			if err != nil {
				return err
			}

			if d != nil {
				// 当前是父进程
				// 父进程直接打印启动成功信息, 不做任何操作
				fmt.Println("cron serve started, pid: ", d.Pid)
				fmt.Println("log file", serverLogFile)
				return nil
			} else {
				// 当前是子进程
				// 子进程执行Cron.Run
				defer ctx.Release() // todo: 研究下release
				fmt.Println("daemon started")
				//gspt.SetProcTitle("hade cron") //感觉像设置进程名称, 应该是调用c语言了
				cmd.Root().Cron.Run()
				return nil
			}

		} else {
			// not daemon mode
			fmt.Println("not daemon mode: start cron job")

			// write pid file
			content := strconv.Itoa(os.Getpid())
			fmt.Println("[PID]", content)
			err := ioutil.WriteFile(serverPidFile, []byte(content), 0664)
			if err != nil {
				return err
			}

			//gspt.SetProcTitle("hade cron")
			log.Printf("cmd is %p\n", cmd)
			log.Printf("cmd.Root() is %p\n", cmd.Root())
			log.Printf("cmd.Root().Cron is %p\n", cmd.Root().Cron)
			// todo: 这块得加个保护, 如果没有设置Cron，那么 cmd.Root().Cron 就是nil，那么下面这么调用就会panic了
			cmd.Root().Cron.Run()
			return nil
		}
	},
}

// list cron task
var cronListCommand = &cobra.Command{
	Use:   "list",
	Short: "列出所有的定时任务",
	RunE: func(cmd *cobra.Command, args []string) error {
		cronSpecs := cmd.Root().CronSpecs
		ps := [][]string{}
		for _, cronSpec := range cronSpecs {
			line := []string{cronSpec.Type, cronSpec.Spec, cronSpec.Cmd.Use, cronSpec.Cmd.Short, cronSpec.ServiceName}
			ps = append(ps, line)
		}
		util.PrettyPrint(ps)

		return nil
	},
}

// todo: cronRestartCommand
var cronRestartCommand = &cobra.Command{
	Use:   "restart",
	Short: "重启cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		// GetPid
		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		content, err := ioutil.ReadFile(serverPidFile)
		if err != nil {
			return err
		}

		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}

			if util.CheckProcessExist(pid) {
				if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
					return nil
				}

				// check process closed
				for i := 0; i < 10; i++ {
					if util.CheckProcessExist(pid) == false {
						break
					}
					time.Sleep(1 * time.Second)
				}
				fmt.Println("kill process:" + strconv.Itoa(pid))
			}
		}

		cronDaemon = true
		// 执行cronStartCommand，且，是daemon模式
		return cronStartCommand.RunE(cmd, args)
	},
}

// todo: cronStopCommand
var cronStopCommand = &cobra.Command{
	Use:   "stop",
	Short: "停止cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		// GetPid
		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		content, err := ioutil.ReadFile(serverPidFile)
		if err != nil {
			return err
		}

		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}

			if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
				return err
			}
			if err := ioutil.WriteFile(serverPidFile, []byte{}, 0644); err != nil {
				return err
			}
			fmt.Println("stop pid: ", pid)
		}
		return nil
	},
}

// todo: cronStateCommand
var cronStateCommand = &cobra.Command{
	Use:   "state",
	Short: "cron常驻进程状态",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		// GetPid
		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		content, err := ioutil.ReadFile(serverPidFile)
		if err != nil {
			return err
		}

		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}

			if util.CheckProcessExist(pid) {
				fmt.Println("cron server started, pid: ", pid)
				return nil
			}
		}
		fmt.Println("no cron server start")
		return nil
	},
}
