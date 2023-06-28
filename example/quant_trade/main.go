package main

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")

	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}

func main() {
	// 1. read conf

	// 2. init

	// 3. go listener

	// 4. 监听signal

	// 5. print 统计信息
}
