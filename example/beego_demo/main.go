package main

import (
	"context"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
)

func main() {
	// create a task
	tk1 := task.NewTask("tk1", "0/3 * * * * *", func(ctx context.Context) error {
		logs.Info("tk1")
		return nil
	})

	// check task
	err := tk1.Run(context.Background())
	if err != nil {
		logs.Error(err)
	}

	// add task to global todolist
	task.AddTask("tk1", tk1)

	// start tasks
	task.StartTask()

	// wait 12 second
	time.Sleep(12 * time.Second)
	defer task.StopTask()
}
