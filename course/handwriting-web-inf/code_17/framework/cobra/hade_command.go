package cobra

import (
	"go-examples/course/handwriting-web-inf/code_17/framework"
	"log"

	"github.com/robfig/cron/v3"
)

// SetContainer 设置服务容器
func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

// GetContainer 获取容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}

// -- section begin: cron
// CronSpec 保存Cron命令的信息，用于展示
type CronSpec struct {
	Type        string
	Cmd         *Command
	Spec        string
	ServiceName string
}

// 设置父节点为null，提升查询效率，不然cobra的command是从root开始遍历到目标节点
func (c *Command) SetParentNull() {
	c.parent = nil
}

// AddCronCommand 用来创建一个Cron任务
func (c *Command) AddCronCommand(spec string, cmd *Command) {
	// cron结构是挂载在根（root）Command上的
	root := c.Root()
	if root.Cron == nil {
		// 初始化cron
		root.Cron = cron.New(cron.WithParser(
			cron.NewParser(
				cron.SecondOptional |
					cron.Minute |
					cron.Hour |
					cron.Dom | // Day of month field, default *
					cron.Month |
					cron.Dow | // Day of week field, default *
					cron.Descriptor))) // Allow descriptors such as @monthly, @weekly, etc.
		// 说明文档 specification
		root.CronSpecs = []CronSpec{}
	}

	// 增加说明信息
	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Type: "normal-cron",
		Cmd:  cmd,
		Spec: spec,
	})

	// 制作一个rootCommand
	var cronCmd Command
	ctx := root.Context()
	cronCmd = *cmd // 得到一个副本
	cronCmd.args = []string{}
	cronCmd.SetParentNull()
	cronCmd.SetContainer(root.GetContainer())

	// 增加调用函数
	entryID, err := root.Cron.AddFunc(spec, func() {
		// 如果后续的command出现panic, 这里要捕获
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		// 这里传进去的是 root.Context()
		// todo: 这个函数还是要回顾下
		err := cronCmd.ExecuteContext(ctx)
		if err != nil {
			// 打印出err信息
			log.Println(err)
		}
	})
	log.Printf("after root.Cron.AddFunc: entryID(%d), err=%s", entryID, err)
}

// -- section end: cron
