package settings

import (
	"fmt"

	fsnotify2 "github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() error {
	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return err
	}
	// 实时监控配置文件的变化
	viper.WatchConfig()
	// 当配置变化之后调用的一个回调函数
	viper.OnConfigChange(func(in fsnotify2.Event) {
		fmt.Println("配置文件修改了...")
	})

	return nil
}
