/**
 * @Author: vincent
 * @Description:
 * @File:  console
 * @Version: 1.0.0
 * @Date: 2021/11/2 17:50
 */

package services

import (
	"go-examples/course/handwriting-web-inf/code_18/framework"
	"go-examples/course/handwriting-web-inf/code_18/framework/contract"
	"os"
)

type HadeConsoleLog struct {
	HadeLog
}

func NewHadeConsoleLog(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &HadeConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = container

	return log, nil
}
