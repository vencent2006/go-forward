/**
 * @Author: vincent
 * @Description:
 * @File:  rotate
 * @Version: 1.0.0
 * @Date: 2021/11/2 17:51
 */

package services

import (
	"fmt"
	"go-examples/course/handwriting-web-inf/code_19/framework"
	"go-examples/course/handwriting-web-inf/code_19/framework/contract"
	"go-examples/course/handwriting-web-inf/code_19/framework/util"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

type HadeRotateLog struct {
	HadeLog

	// 日志文件存储目录
	folder string

	// 日志文件名
	file string
}

func NewHadeRotateLog(params ...interface{}) (interface{}, error) {
	// 参数解析
	container := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	appService := container.MustMake(contract.AppKey).(contract.App)
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// 从配置文件中获取folder信息，否则使用默认的LogFolder文件夹
	folder := appService.LogFolder()
	if configService.IsExist("log.folder") {
		folder = configService.GetString("log.folder")
	}

	// 如果folder不存在，则创建
	if !util.Exists(folder) {
		// todo: 这个返回值还是应该接一下的
		os.MkdirAll(folder, os.ModePerm)
	}

	// 从配置文件中获取file信息，否则使用默认的hade.log
	file := "hade.log"
	if configService.IsExist("log.file") {
		file = configService.GetString("log.file")
	}

	// 从配置文件获取date_format信息
	dateFormat := "%Y%m%d%H"
	if configService.IsExist("log.date_format") {
		dateFormat = configService.GetString("log.date_format")
	}

	// ---- rotate
	linkName := rotatelogs.WithLinkName(filepath.Join(folder, file))
	options := []rotatelogs.Option{linkName}

	// 从配置文件获取 rotate_count 信息
	if configService.IsExist("log.rotate_count") {
		rotateCount := configService.GetInt("log.rotate_count")
		options = append(options, rotatelogs.WithRotationCount(uint(rotateCount)))
	}

	// 从配置文件获取 rotate_size 信息
	if configService.IsExist("log.rotate_size") {
		rotateSize := configService.GetInt("log.rotate_size")
		options = append(options, rotatelogs.WithRotationSize(int64(rotateSize)))
	}

	// 从配置文件获取max_age信息
	if configService.IsExist("log.max_age") {
		if maxAgeParse, err := time.ParseDuration(configService.GetString("log.max_age")); err == nil {
			options = append(options, rotatelogs.WithMaxAge(maxAgeParse))
		}
	}

	// 从配置文件获取rotate_time信息
	if configService.IsExist("log.rotate_time") {
		if rotateTimeParse, err := time.ParseDuration(configService.GetString("log.rotate_time")); err == nil {
			options = append(options, rotatelogs.WithRotationTime(rotateTimeParse))
		}
	}

	// 设置基础信息
	log := &HadeRotateLog{}
	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)
	log.folder = folder
	log.file = file

	w, err := rotatelogs.New(fmt.Sprintf("%s.%s", filepath.Join(log.folder, log.file), dateFormat), options...)
	if err != nil {
		// todo: 这里使用了errors.Wrap
		return nil, errors.Wrap(err, "new rotatelogs error")
	}
	log.SetOutput(w)
	log.c = container

	return log, nil
}
